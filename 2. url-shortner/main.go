package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	urlshort "url-shortner/mapHandler"
)

type Config struct {
	pathToJSON string
	pathToYAML string
	// pathToBoltDB string
}

func main() {
	config := getConfig()

	mux := makeDefaultMux()

	yamlBytes := getFileBytes(config.pathToYAML)
	jsonBytes := getFileBytes(config.pathToJSON)

	makeMapHandler(mux)
	makeYAMLHandler(yamlBytes, mux)
	makeJSONHandler(jsonBytes, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}

func getConfig() *Config {
	config := Config{}

	config.pathToJSON = *flag.String("json-path", "./storage/url.json", "file location of json containing url redirections")
	config.pathToYAML = *flag.String("yaml-path", "./storage/url.yaml", "file location of yaml containing url redirections")
	// config.pathToBoltDB = *flag.String("boltdb-path", "./storage/bolt.db", "file location of boltdb containing url redirections")
	flag.Parse()
	return &config
}

func getFileBytes(path string) []byte {
	bytes, err := ioutil.ReadFile(path) // better than opening a file since you dont have to take care about closing that givern file too
	if err != nil {
		return nil
	}
	return bytes
}

func indexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func makeDefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexController)
	return mux
}

func makeMapHandler(mux *http.ServeMux) {
	pathsToUrls := map[string]string{
		"/lin": "https://www.linkedin.com/in/rishabh-malhotra-4536a418b/",
		"/mt":  "https://monkeytype.com/",
	}
	urlshort.MapHandler(pathsToUrls, mux)
}

func makeYAMLHandler(yamlBytes []byte, fallbackHandler *http.ServeMux) {
	err := urlshort.YAMLHandler(yamlBytes, fallbackHandler)
	if err != nil {
		panic(err)
	}
}

func makeJSONHandler(jsonBytes []byte, fallbackHandler *http.ServeMux) {
	err := urlshort.JSONHandler(jsonBytes, fallbackHandler)
	if err != nil {
		panic(err)
	}
}
