package models

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Phone struct {
	Id int 		`db:"id"`
	Name string	`db:"name"`	
	Phone string `db:"phone"`
}

// func UpdatePhones(phones []Phone) {
// 	statement :=  s

// }

func GetPhone(phone string) (Phone, error) {
	p := Phone{}
	err := db.Get(&p, "SELECT * FROM phone WHERE phone=$1", phone)
	if err != nil {
		return Phone{}, fmt.Errorf("falied to get phone: %v", err)
	}
	return p, nil
}

func UpdatePhone(p Phone, norm string) error{
	_, err := db.Exec("UPDATE phone SET value=$2 WHERE id=$1", p.Id, norm)
	if err != nil {
		return fmt.Errorf("failed to update %v: %v", p, err)
	}
	return nil
}

func DeletePhone(phone string) error {
	_,err := db.NamedExec("DELETE FROM phone where phone=:phone", phone)
	if err != nil {
		return fmt.Errorf("failed to delete %v: %v", phone, err)
	}
	return nil
}

func GetAllData() ([]Phone, error) {
	phones := []Phone{}
	err := db.Select(&phones, "SELECT * FROM phone")
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %v", err)
	}
	return phones, nil
}

func InitDB(databaseInfo string) error {
	var err error
	db, err = sqlx.Connect("postgres", databaseInfo)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	schema := `CREATE TABLE IF NOT EXISTS phone (
		"id" SERIAL,
		"name" VARCHAR(50),
		"phone" VARCHAR(20),
		PRIMARY KEY (id)
	);`

	db.MustExec(schema)
	log.Println("Phone table created.")
	return db.Ping()
}

func AddToDB(phones []string) error {

	tx := db.MustBegin()
	for _, phone := range phones {
		_, err := tx.NamedExec("INSERT INTO phone (name, phone) VALUES (:name, :phone)", Phone{Name: randString(20), Phone: phone})
		if err != nil {
			return fmt.Errorf("failed to insert person %v: %v", phone, err)
		}
	}
	tx.Commit()

	return nil
}

func randString(n int) string {
	// 長さnの任意の文字列を作る
	rand.Seed(time.Now().UnixNano())
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
