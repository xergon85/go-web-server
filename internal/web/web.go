package web

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() error {

	router, err := createRouter()
	if err != nil {
		return err
	}

	if err := http.ListenAndServe(":80", router); err != nil {
		return err
	}

	return nil
}

func createRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	handleStatic(r)
	r.HandleFunc("/", indexGet)
	r.HandleFunc("/books/{title}/page/{page}", bookTitlePage)

	return r, nil
}

func indexGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helllo, you've requested: %s\n", r.URL.Path)
}

func handleStatic(r *mux.Router) error {
	var dir string

	flag.StringVar(&dir, dir, "static", "the directory to serve files from. Defaults to the /static")
	flag.Parse()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	return nil
}

func bookTitlePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}
