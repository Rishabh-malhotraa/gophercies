package maphandler

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type urlMap struct {
	cache map[string]string
}

type MapItems struct {
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

func MapHandler(items []MapItems, mux *http.ServeMux) {
	for _, item := range items {
		globalMap.cache[item.Path] = item.Url
		mux.HandleFunc(item.Path, globalMap.redirect)
	}
}

func YAMLHandler(buffer []byte, mux *http.ServeMux) error {
	var items []MapItems
	err := yaml.Unmarshal(buffer, &items)

	if err != nil {
		return err
	}

	MapHandler(items, mux)
	return nil
}

func JSONHandler(buffer []byte, mux *http.ServeMux) error {
	var items []MapItems
	err := json.Unmarshal(buffer, &items)

	if err != nil {
		return err
	}

	MapHandler(items, mux)
	return nil
}
