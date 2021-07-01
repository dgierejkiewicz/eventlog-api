package db

import (
		"database/sql"
		"log"
)

var DB *sql.DB

// InitDB Function responsible for connection initialization to db
func InitDB(dataSourceName string) {

		var err error
		DB, err = sql.Open("postgres", dataSourceName)
		if err != nil {
				log.Println(err)
		}

		if err = DB.Ping(); err != nil {
				log.Println(err)
		} else {
				log.Println("Connection to db established!")
		}

}
