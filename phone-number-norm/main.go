package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/smelton01/go-basics/phone-number-norm/models"
)

type server struct {
	db *sql.DB
	psqlInfo string
}



func main() {
	databaseInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "simonjd", "postgres")
	if err := models.InitDB(databaseInfo); err != nil {
		panic(err)
	}

	file, err := os.Open("phones.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var phones []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phones = append(phones, scanner.Text())
	}

	if err := models.AddToDB(phones); err!=nil {
		panic(err)
	}
}

