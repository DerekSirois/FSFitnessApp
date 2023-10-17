package database

import "fmt"

type User struct {
	Id       int
	Username string
	Email    string
	Password []byte
}

func (u *User) Create() error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", u.Username, u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("couldn't create the user: %v", err)
	}
	return nil
}

func (u *User) GetByUsername(username string) error {
	err := db.Get(u, "SELECT * FROM users WHERE username = $1", username)
	return err
}
