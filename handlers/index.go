package handlers

import (
	"net/http"
	"sync"
	"os"
	"kenef.online/dep"
)

func index(w http.ResponseWriter, r *http.Request, d *dep.Dependencies, wg *sync.WaitGroup) {
	defer wg.Done()

	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// get home dir
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", 500)
	// 	d.Log.Error(err)
	// 	return
	// }

	// read music dir
	files, err := os.ReadDir("./music/")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// make track list
	var trackList []string
	for _, file := range files {
		trackList = append(trackList, file.Name())
	}

	// return template
	if err := d.Tmp.ExecuteTemplate(w, "index.html", trackList); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
	}
}