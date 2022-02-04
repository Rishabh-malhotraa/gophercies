package maphandler

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type urlMap struct {
	cache map[string]string
}

type mapItems struct {
	Path string
	Url  string
}

var globalMap *urlMap = &urlMap{make(map[string]string)}

func (*urlMap) redirect(w http.ResponseWriter, r *http.Request) {
	if redirectPath, isPresent := globalMap.cache[r.URL.Path]; isPresent == true {
		http.Redirect(w, r, redirectPath, 307)
		return
	}
}

func MapHandler(pathsToUrls map[string]string, mux *http.ServeMux) {
	for k, v := range pathsToUrls {
		globalMap.cache[k] = v
		mux.HandleFunc(k, globalMap.redirect)
	}
}

func YAMLHandler(yml []byte, mux *http.ServeMux) error {
	var items []mapItems
	err := yaml.Unmarshal(yml, &items)

	if err != nil {
		return err
	}

	for _, item := range items {
		globalMap.cache[item.Path] = item.Url
		mux.HandleFunc(item.Path, globalMap.redirect)
	}

	return nil
}

func JSONHandler(buffer []byte, mux *http.ServeMux) error {
	var items []mapItems
	err := json.Unmarshal(buffer, &items)

	if err != nil {
		return err
	}

	for _, item := range items {
		globalMap.cache[item.Path] = item.Url
		mux.HandleFunc(item.Path, globalMap.redirect)
	}

	return nil
}
