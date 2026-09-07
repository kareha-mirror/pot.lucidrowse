package api

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/unicode/norm"

	"tea.kareha.org/pot/lucidrowse/server/internal/config"
	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
	return err == nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func newSessionWithUserID(
	cfg *config.Config, w http.ResponseWriter, userID int,
) {
	keyHash, err := newSessionKeyHash(cfg, w)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"セッションキーを作れません。",
		)
		return
	}

	if err = data.AddSession(userID, keyHash); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusInternalServerError,
			"セッションを作れません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}

func (api *API) login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
		)
		return
	}

	username := norm.NFC.String(req.Username)
	password := norm.NFC.String(req.Password)

	user, err := data.LoadUserByName(username)
	if err != nil {
		writeError(
			w,
			http.StatusUnauthorized,
			"鍵がありません。",
		)
		return
	}

	if !CheckPassword(*user.PassHash, password) {
		writeError(
			w,
			http.StatusUnauthorized,
			"合い言葉が違います。",
		)
		return
	}

	newSessionWithUserID(api.cfg, w, user.ID)
}

type CreateKeyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *API) createKey(w http.ResponseWriter, r *http.Request) {
	user, err := auth(w, r)
	if err != nil {
		return
	}

	var req CreateKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
		)
		return
	}

	username := norm.NFC.String(req.Username)
	password := norm.NFC.String(req.Password)

	hash, err := HashPassword(password)
	if err != nil {
		writeError(
			w,
			http.StatusInternalServerError,
			"合い言葉を作れません。",
		)
		return
	}

	err = data.CreateKey(user.ID, username, hash)
	if err != nil {
		writeError(
			w,
			http.StatusInternalServerError,
			"同じ名前の鍵があるので作れません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new-password"`
}

func (api *API) changePassword(w http.ResponseWriter, r *http.Request) {
	user, err := auth(w, r)
	if err != nil {
		return
	}

	var req ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusBadRequest,
			"変な要求です。",
		)
		return
	}

	if user.Name == nil {
		writeError(
			w,
			http.StatusBadRequest,
			"まだ鍵を作ってません。",
		)
		return
	}

	username := *user.Name
	password := norm.NFC.String(req.Password)
	newPassword := norm.NFC.String(req.NewPassword)

	user, err = data.LoadUserByName(username)
	if err != nil {
		writeError(
			w,
			http.StatusUnauthorized,
			"鍵がありません。",
		)
		return
	}

	if !CheckPassword(*user.PassHash, password) {
		writeError(
			w,
			http.StatusUnauthorized,
			"合い言葉が違います。",
		)
		return
	}

	newHash, err := HashPassword(newPassword)
	if err != nil {
		writeError(
			w,
			http.StatusInternalServerError,
			"合い言葉を作れません。",
		)
		return
	}

	err = data.ChangePassword(user.ID, newHash)
	if err != nil {
		writeError(
			w,
			http.StatusInternalServerError,
			"合い言葉を変えられません。",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}

func (api *API) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"秘密の鍵を使ってません。",
		)
		return
	}

	sessionKey := cookie.Value
	keyHash := sha256.Sum256([]byte(sessionKey))
	err = data.RevokeSession(keyHash[:])
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"秘密の鍵を外せません。",
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   api.cfg.App.Mode == "release",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{}{})
}
