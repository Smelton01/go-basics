package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/smelton01/go-basics/urlshort"
)

const dbPath = "urlshort.db"

func main() {
	var filePath = flag.String("file", "paths.json", "Path to file with URL shortcuts")
	var initFlag = flag.Bool("newdb", false, "Set true to initialize  anew database in local storage")
	flag.Parse()

	var db *sql.DB
	if *initFlag {
		db = initDB(dbPath)
	}else{
		var err error
		db, err = sql.Open("sqlite3", "./urlshort.db")
		if err != nil {
			log.Fatal("Database read error: ", err)
		}
	}

	mux := defaultMux()

	var handler http.Handler
	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal("File error: ", err)
	}

	defaultHandler := urlshort.DBHandler(db, mux)

	if ext := filepath.Ext(*filePath); ext == ".yaml" {
		handler, err = urlshort.YAMLHandler(file, defaultHandler)

		if err != nil {
			log.Fatal("Handler Error", err)
		}
	} else if ext == ".json" {
		handler, err = urlshort.JSONHandler(file, defaultHandler)
		if err != nil {
			log.Fatal("Handler Error", err)
		}
	}else {
		handler = mux
		// fmt.Println(db)
	}
	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func initDB(path  string) *sql.DB {

	err := os.Remove(path)
	if err != nil {
		log.Println("Database not found.")
	}
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	createURLShortTable := `CREATE TABLE urlshort (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"path" TEXT,
		"url" TEXT 
	);`

	stmt, err := db.Prepare(createURLShortTable)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	log.Println("URL Shorterner database created.")
	return db
}