package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/smelton01/go-basics/urlshort"
)

type URLShort struct{
	Path string
	URL string
}

type dbObject struct{
	DB *sql.DB
}

var Database dbObject

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

	Database.DB = db
	index := defaultMux()

	var handler http.Handler
	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal("File error: ", err)
	}

	defaultHandler := urlshort.DBHandler(db, index)

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
		handler = index
	}
	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Database.hello)
	return mux
}

func (db *dbObject) hello(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/layout.html"))

    if r.Method != http.MethodPost {
        tmpl.Execute(w, nil)
        return
    }
    details := URLShort{
        Path:   r.FormValue("path"),
        URL: r.FormValue("url"),
    }
    // do something with details
    
	err := db.addURL(details.Path, details.URL)
	if err != nil {
		log.Fatal("Insert Error: ", err)
	}

    tmpl.Execute(w, struct{ 
		Success bool;
		URL string
		}{true, fmt.Sprintf("{ /%v : => :%v }", details.Path,details.URL)})
	fmt.Fprintln(w, "Hello, world!", details.Path)
}

func (db *dbObject) addURL(path, url string) error {
	insertURL := `INSERT INTO urlshort(path, url) values(?, ?)`
	stmt, err := db.DB.Prepare(insertURL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec("/"+path, url)
	if err != nil {
		return err
	}
	return nil
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