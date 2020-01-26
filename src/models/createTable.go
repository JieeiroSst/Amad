package models

import (
	"database/sql"
	"keikibook/src/dto"
	"log"
)

var db *sql.DB
var user dto.User

func CreateTable() {
	query := `CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		PRIMARY KEY (id)
	    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertTableUser(FirstName, LastName, CompayName, Email, Contry, Address, Town, ZipCode, Phone, Note string) {
	_, err := db.Exec(`INSERT INTO User VALUES (?,?,?,?,?,?,?,?,?,?)`, FirstName, LastName, CompayName, Email, Contry, Address, Town, ZipCode, Phone, Note)
	if err != nil {
		log.Fatal(err)
	}
}
