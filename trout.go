package main

import (
    "html/template"
    "log"
    "net/http"
    "regexp"
)

var templates = template.Must(template.ParseFiles("trout.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.Execute(w, struct{}{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var staticPath = regexp.MustCompile("^/static/([a-zA-Z0-9_.]+)$")

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	fileName := staticPath.FindStringSubmatch(r.URL.Path)
	if fileName == nil || len(fileName) <= 1 || (fileName[1] != "cat.jpg" && fileName[1] != "favicon.ico"){
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, fileName[1])
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/static/", StaticHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
