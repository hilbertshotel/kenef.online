package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"kenef.online/dep"
)

func index(w http.ResponseWriter, r *http.Request, d *dep.Dependencies, wg *sync.WaitGroup) {
	defer wg.Done()

	// handle GET method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// read file dir
	files, err := os.ReadDir(d.Cfg.FileDir)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}

	// make file list & file map
	var fileList []string
	fileMap := make(map[string]string)

	for _, file := range files {
		fileList = append(fileList, file.Name())

		fullPath := filepath.Join(d.Cfg.FileDir, file.Name())
		data, err := os.ReadFile(fullPath)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}

		fileMap[file.Name()] = string(data)
	}

	// check url
	urlPath := strings.TrimPrefix(r.URL.Path, "/")

	// serve index template
	if urlPath == "" {
		if err := d.Tmp.ExecuteTemplate(w, "index.html", fileList); err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
		}
		return
	}

	// serve story template
	if err := d.Tmp.ExecuteTemplate(w, "story.html", [2]string{urlPath, fileMap[urlPath]}); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
	}
}
