package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type path2URL struct {
	Path string `yaml:"path"`
	Url  string	`yaml:"url"`
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
