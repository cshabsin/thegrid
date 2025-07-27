package main

import (
	"archive/zip"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var registeredApps []string

func registerApp(name, zipPath string) {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Printf("failed to open %s: %v. Skipping.", zipPath, err)
		return
	}
	http.Handle(fmt.Sprintf("/%s/", name), http.StripPrefix(fmt.Sprintf("/%s/", name), http.FileServer(http.FS(zipReader))))
	registeredApps = append(registeredApps, name)
	log.Printf("Registered app '%s' from %s", name, zipPath)
}

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "_pkg.zip") {
			appName := strings.TrimSuffix(info.Name(), "_pkg.zip")
			registerApp(appName, path)
		}
		return nil
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("index").Parse(`
			<html>
				<head>
					<title>The Grid</title>
				</head>
				<body>
					<h1>Available Services</h1>
					<ul>
						{{range .}}
						<li><a href="/{{.}}">{{.}}</a></li>
						{{end}}
					</ul>
				</body>
			</html>
		`)
		if err != nil {
			log.Printf("template parse error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := t.Execute(w, registeredApps); err != nil {
			log.Printf("template execute error: %v", err)
		}
	})

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}