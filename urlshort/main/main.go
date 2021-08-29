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
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/smelton01/go-basics/urlshort"
)

type URLShort struct{
	Path string
	URL string
	Timeout string
}

type dbObject struct{
	DB *sql.DB
}

var Database dbObject

const dbPath = "urlshort.db"

func main() {
	var filePath = flag.String("file", "./paths.json", "Path to file with URL shortcuts")
	var initFlag = flag.Bool("newdb", false, "Set true to initialize  anew database in local storage")
	// var timeout = flag.Int("time", 24, "Timeout in hours of created URL shortcut")
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
	mux := defaultMux()

	var handler http.Handler
	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Println("File error: ", err)
		file = []byte(`[{"path":"/test", "url":"www.example.com"}]`)
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
	}
	
	go func () {
		for {
			Database.cleanup()
			time.Sleep(time.Duration(10)*time.Minute)
		}
	}()
		

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Database.index)
	return mux
}

func (db *dbObject) index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/layout.html"))

    if r.Method != http.MethodPost {
        tmpl.Execute(w, nil)
        return
    }
    details := URLShort{
        Path:   r.FormValue("path"),
        URL: r.FormValue("url"),
    }
    timeout, err :=  strconv.Atoi(r.FormValue("timeout"))
	if err != nil {
		log.Fatal("Invalid timeout")
	}
	exp := time.Now().Add(time.Duration(timeout)*time.Minute).Format(time.RFC3339)

	err = db.addURL(details.Path, details.URL, exp)
	if err != nil {
		log.Fatal("Insert Error: ", err)
	}

    tmpl.Execute(w, struct{ 
		Success bool;
		URL string
		}{true, fmt.Sprintf("{ /%v : => :%v }", details.Path,details.URL)})
	fmt.Fprintln(w, "Expires at :", func() string {
		str, err := time.Parse(time.RFC3339, exp)
		if err != nil {
			log.Fatal(err)
		}
		return str.Format(time.RFC1123)
	}())
}

func (db *dbObject) addURL(path, url string, timeout string) error {
	insertURL := `INSERT INTO urlshort(path, url, timeout) values(?, ?, ?)`
	stmt, err := db.DB.Prepare(insertURL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec("/"+path, url, timeout)
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
		log.Fatal("File create: ",err)
	}
	defer file.Close()
	
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	createURLShortTable := `CREATE TABLE urlshort (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"path" TEXT,
		"url" TEXT,
		"timeout" STRING 
	);`

	stmt, err := db.Prepare(createURLShortTable)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	log.Println("URL Shorterner database created.")
	return db
}

func (db *dbObject) cleanup() {
	expiredURLs := `SELECT * FROM urlshort where timeout < ?`

	rows, err := db.DB.Query(expiredURLs, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println("Clean error", err)
	}
	defer rows.Close()
	var trash []struct{
		path string
	}
	var id int
	var path, url, timeout string
	for rows.Next() {
		err := rows.Scan(&id, &path, &url, &timeout)
		if err != nil {
			log.Println(err)
			return
		}
		trash = append(trash, struct{path string}{path})
	}
	rows.Close()

	for item := range trash {
		deleteEntry := `DELETE from urlshort where path=?`
		_ , err = db.DB.Exec(deleteEntry, item)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(path, "URL deleted successfully")
	}
	
}