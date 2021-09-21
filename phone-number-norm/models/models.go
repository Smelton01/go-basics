package models

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Phone struct {
	id int
	name string
	phone string
}

func InitDB(databaseInfo string) error {
	var err error
	db, err = sqlx.Connect("postgres", databaseInfo)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	createCardsTable := `CREATE TABLE IF NOT EXISTS phone (
		"id" SERIAL,
		"name" VARCHAR(50),
		"phone" VARCHAR(20),
		PRIMARY KEY (id)
	);`

	stmt, err := db.Prepare(createCardsTable)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %s", err)
	}
	stmt.Exec()
	log.Println("Phone table created.")
	return db.Ping()
}

func AddToDB(phones []string) error {
	values := []interface{}{}
	statement := "INSERT INTO phone (name, phone) VALUES"
	// build statement for bulk insert
	for i, num := range phones {
		values = append(values, randString(20))
		values = append(values, num) 
		statement += fmt.Sprintf(" ($%v, $%v),", i*2+1, i*2+2)
	}
	stmt, err := db.Prepare(statement)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %s", err)
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}
	return nil
}

func randString(n int) string {
	// 長さnの任意の文字列を作る
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
