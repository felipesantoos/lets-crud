package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() (*sql.DB, error) {
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

func closeConnection(conn *sql.DB) error {
	err := conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetConnection() (*sql.DB, error) {
	return getConnection()
}
