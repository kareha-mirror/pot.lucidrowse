package data

func Seed() error {
	err := SeedAreas()
	if err != nil {
		return err
	}
	return nil
}
