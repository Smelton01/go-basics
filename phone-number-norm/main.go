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
	p, err := models.GetAllData() 
	if err != nil {
		panic(err)
	}
	for _, phone := range p {
		normPhone := normalize(phone.Phone)
		if normPhone == phone.Phone {
			fmt.Printf("No updates needed")
			continue
		}
		existing, err := models.GetPhone(normPhone)
		if err != nil {
			panic(err)
		}
		n := models.Phone{}
		if existing != n {
			err := models.DeletePhone(phone.Phone)
			if err != nil {
				panic(err)
			}
		}else{
			models.UpdatePhone(phone, normPhone)
		}
	}
}

func normalize(number string) string {
	var normalForm = ""
	for _, c := range number {
		if c < rune('1') || c > rune('9'){
			continue
		}
		normalForm += string(c)
	}
	return normalForm
}

