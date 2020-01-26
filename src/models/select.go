package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SelectUserId(firstName, lastname, compayName, email, contry, address, town, zipcode, phone, note string) {
	query := `SELECT * FROM Users where id=?`
	err := db.QueryRow(query, 1).Scan(&firstName, &lastname, &compayName, &email, &contry, &address, &town, &zipcode, &phone, &note)
	if err != nil {
		log.Fatal(err)
	}
}

func SelectAllUser(firstName, lastname, compayName, email, contry, address, town, zipcode, phone, note string) {
	rows, err := db.Query(`SELECT * FROM Users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if rows.Err()!=nil{ return}
		err := rows.Scan(&firstName, &lastname, &compayName, &email, &contry, &address, &town, &zipcode, &phone, &note)
		if err != nil {
			log.Fatal(err)
		}
	}
}
