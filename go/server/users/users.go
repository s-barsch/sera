package users

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	db *sql.DB
}

type User struct {
	Name     string
	Mail     string
	PaypalID string `json:"subscriptionID"`
}

func LoadUsers() (*Users, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}
	return &Users{db: db}, nil
}

func (u *Users) Close() error {
	return u.db.Close()
}

func (u *Users) StoreSession(mail, key string) error {
	_, err := u.db.Exec(
		"INSERT INTO session(mail, key) VALUES(?,?);",
		mail,
		key,
	)
	return err
}

func (u *Users) StoreVerify(mail, key string) error {
	_, err := u.db.Exec(
		"INSERT INTO verify(mail, key) VALUES(?,?);",
		mail,
		key,
	)
	return err
}

func (u *Users) GetSessionKey(mail string) (string, error) {
	var str string

	return str, u.db.QueryRow("SELECT key FROM session WHERE mail=?", mail).Scan(&str)
}

func (u *Users) GetVerify(mail string) (string, error) {
	var str string

	return str, u.db.QueryRow("SELECT key FROM verify WHERE mail=?", mail).Scan(&str)
}

func (u *Users) AddUser(user *User) error {
	_, err := u.db.Exec(
		"INSERT INTO users(name, mail) VALUES(?,?);",
		user.Name,
		user.Mail,
	)
	if err != nil {
		return fmt.Errorf("insert: %v", err)
	}
	return nil
}

func (u *Users) LookupUser(mail string) (*User, error) {
	user := &User{}

	return user, u.db.QueryRow("SELECT name, mail, paypal_id FROM users WHERE mail=?", mail).Scan(&user.Name, &user.Mail, &user.PaypalID)
}
