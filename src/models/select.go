package models

import (
	"log"
)

func SelectUserId(firstName, lastname, compayName, email, contry, address, town, zipcode, phone, note string) {
	query := `SELECT * FROM users where id=?`
	err := db.QueryRow(query, 1).Scan(&firstName, &lastname, &compayName, &email, &contry, &address, &town, &zipcode, &phone, &note)
	if err != nil {
		log.Fatal(err)
	}
}

func SelectAllUser(firstName, lastname, compayName, email, contry, address, town, zipcode, phone, note string) {
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if rows.Err() != nil {
			return
		}
		err := rows.Scan(&firstName, &lastname, &compayName, &email, &contry, &address, &town, &zipcode, &phone, &note)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SelectAllProduct(name, price, note, img string) {
	rows, err := db.Query(`SELECT * FROM product`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if rows.Err() != nil {
			return
		}
		err := rows.Scan(&name, &price, &note, &img)
		if err != nil {
			log.Fatal(err)
		}
	}
}
