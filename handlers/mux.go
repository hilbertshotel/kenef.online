package handlers

import (
	"net/http"
	"sync"

	"kenef.online/dep"
)

func Mux(d *dep.Dependencies, wg *sync.WaitGroup) *http.ServeMux {
	mux := http.NewServeMux()

	// static files
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	// index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		index(w, r, d, wg)
	})

	return mux
}
