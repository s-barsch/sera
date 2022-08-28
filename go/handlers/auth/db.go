package auth

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func storeSession(mail, key string) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO session(mail, key) VALUES(?,?);",
		mail,
		key,
	)
	return err
}

func storeVerifyKey(mail, key string) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO verify(mail, key) VALUES(?,?);",
		mail,
		key,
	)
	return err
}

func getVerifyKey(mail string) (string, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var str string

	return str, db.QueryRow("SELECT key FROM verify WHERE mail=?", mail).Scan(&str)
}



func add(u *User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users(name, mail) VALUES(?,?);",
		u.Name,
		u.Mail,
	)
	if err != nil {
		return fmt.Errorf("insert: %v", err)
	}
	return nil
}

func lookup(mail string) (*User, error) {
	user := &User{}

	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return user, db.QueryRow("SELECT name, mail FROM users WHERE mail=?", mail).Scan(&user.Name, &user.Mail)
}
