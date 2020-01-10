package models

import (
	"database/sql"
	"fmt"
	"keikibook/src/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Querying a Single Row
//var tag Tag
// Execute the query
//err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
//if err != nil {
//   panic(err.Error()) // proper error handling instead of panic in your app
//}
//log.Println(tag.ID)
//log.Println(tag.Name)

func init() {
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open(config.Server, config.Username+":"+config.Password+"@tcp("+config.Localhost+":"+config.Port+")/"+config.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
