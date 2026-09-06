package api

import (
	"crypto/sha256"
	"log"
	"net/http"

	"tea.kareha.org/pot/lucidrowse/server/internal/data"
)

func auth(w http.ResponseWriter, r *http.Request) (data.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusUnauthorized,
			"あなたが誰なのか分かりません。",
		)
		return data.User{}, err
	}

	keyHash := sha256.Sum256([]byte(cookie.Value))
	user, err := data.LoadUser(keyHash[:])
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusUnauthorized,
			"あなたが誰なのか覚えてません。",
		)
		return data.User{}, err
	}

	return user, nil
}

func authPlayer(
	w http.ResponseWriter, r *http.Request,
) (data.User, data.Player, error) {
	user, err := auth(w, r)
	if err != nil {
		return data.User{}, data.Player{}, err
	}

	player, err := data.LoadPlayer(user.ID)
	if err != nil {
		log.Println(err)
		writeError(
			w,
			http.StatusNotFound,
			"まだこの世界に住んでません。",
		)
		return data.User{}, data.Player{}, err
	}

	return user, player, nil
}
