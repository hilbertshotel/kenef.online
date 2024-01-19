package handlers

import (
	"net/http"
	"kenef.online/dep"
	"sync"
)

func Mux(d *dep.Dependencies, wg sync.WaitGroup) *http.ServeMux {
	mux := http.NewServeMux()

	// static files
	frontend := http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend/")))
	mux.Handle("/frontend/", frontend)

	music := http.StripPrefix("/music/", http.FileServer(http.Dir("./music/")))
	mux.Handle("/music/", music)

	// index
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		index(w, r, d, &wg)
	})

	return mux
}