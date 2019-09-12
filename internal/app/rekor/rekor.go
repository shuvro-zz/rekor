package rekor

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl *template.Template

// RunUI runs the UI
func RunUI() {
	tmpl, err := template.ParseGlob("ui/templates/*.html")
	if err != nil {
		log.Fatalf("unable to parse the templates: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", "MyTitle")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
