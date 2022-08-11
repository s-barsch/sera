package api

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func add(u *User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	res, err := db.Exec(
		"INSERT INTO users(name, mail) VALUES(?,?);",
		u.Name,
		u.Mail,
	)
	if err != nil {
		return fmt.Errorf("insert: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
