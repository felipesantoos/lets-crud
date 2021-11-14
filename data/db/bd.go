package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	user := "root"
	pwd := ""
	name := "letscrud"
	dbURL := fmt.Sprintf("%s:%s@/%s", user, pwd, name)

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Println("Error while acessing database: " + err.Error())

		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
