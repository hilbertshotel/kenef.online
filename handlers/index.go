package handlers

import (
	"net/http"
	"sync"
	"kenef.online/dep"
)

func index(w http.ResponseWriter, r *http.Request, d *dep.Dependencies, wg *sync.WaitGroup) {
	defer wg.Done()

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// return template
	if err := d.Tmp.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
	}
}