package tests

import (
	"letscrud/data/db"
)

const CLEAR_CUSTOMER_TABLE = "DELETE FROM customer WHERE id != 0"
const RESET_AUTO_INCREMET = "ALTER TABLE customer AUTO_INCREMENT = 0"
const INSERT_CUSTOMER_ID_1 = `
	INSERT INTO customer (id, cpf, name, birthDate) VALUES 
	(1, "09862956046", "Lyara Caparica Onofre", "2001-01-01")
`
const INSERT_CUSTOMER_ID_2 = `
	INSERT INTO customer (id, cpf, name, birthDate) VALUES
	(2, "78045161000", "Christopher Graça Sacramento", "2002-02-02")
`

func SetUp(queries []string) {
	conn, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	for _, query := range queries {
		_, err = conn.Exec(query)
		if err != nil {
			panic(err)
		}
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
