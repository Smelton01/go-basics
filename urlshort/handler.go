package urlshort

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"
)

type path2URL struct {
	Path string `yaml:"path" json:"path"`
	Url  string	`yaml:"url" json:"url"`
}

// MapHandler will return an http.HandlerFunc that will attempt to map any
// paths to their corresponding URL 
// If the path is not provided in the map, then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
		if url != "" {
			http.Redirect(res, req, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	})
}

// DBHandler will query the database for provided path
// and redirect the user to the URL corresponding to the shortcut
// If the path is not found in the database, then the
// fallback http.Handler will be called instead.
func DBHandler(db *sql.DB, fallback  http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request)  {
		path := req.URL.Path
		mapDB := dbQuery(db, path)
		if p := mapDB[path]; p != "" {
			http.Redirect(res, req, p, http.StatusTemporaryRedirect)
		}else{
			fallback.ServeHTTP(res, req)
		}
	})
}


func dbQuery(db *sql.DB, path string) map[string]string {
	rows, err := db.Query("SELECT * FROM urlshort where path=?", path)
	if err != nil {
		log.Fatal("Database read error: ", err)
	}
	defer rows.Close()

	out := make(map[string]string)
	for rows.Next() {
		var path, url string 
		var id int
		var timeout string
		if err := rows.Scan(&id, &path, &url, &timeout); err != nil {
			log.Fatal(err)
		}
		exp, err := time.Parse(time.RFC3339, timeout)
		if err != nil {
			log.Fatal("Time error", err)
		}
		if time.Until(exp) < time.Duration(0) {
			fmt.Println("Expired", time.Until(exp))
			return out
		}
		out[path] = url
	}
	return out
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
  	if err != nil {
  	  return nil, err
  	}
  	pathMap := buildMap(parsedYaml)
  	return MapHandler(pathMap, fallback), nil	
}

func parseYAML(yml []byte) ([]path2URL, error) {
	var data []path2URL
	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		fmt.Println("YAML parse error: ", err)
		return []path2URL{}, err
	}
	return data, nil
}

func buildMap(yml []path2URL) map[string]string {
	ymlMap :=  make(map[string]string)
	for _, val := range yml {
		ymlMap[val.Path] = val.Url
	}
	return ymlMap
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(json)
	if err != nil {
		return nil, err
	}
	jsonMap := buildMap(parsedJSON) 
	return MapHandler(jsonMap, fallback), nil
}

func parseJSON(js []byte) ([]path2URL, error) {
	var parsedJSON []path2URL
	err := json.Unmarshal(js, &parsedJSON)
	if err != nil {
		return nil, err
	}
	return parsedJSON, nil
}
