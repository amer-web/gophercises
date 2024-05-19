package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type adventure map[string]story
type story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}
type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	var adventure adventure
	file, _ := os.Open("gopher.json")
	defer file.Close()
	json.NewDecoder(file).Decode(&adventure)

	tmpl := template.Must(template.ParseFiles("layout.html"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		segments := strings.TrimSpace(r.URL.Path)[1:]
		data, ok := adventure[segments]
		if !ok {
			data = adventure["intro"]
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", mux)
}
