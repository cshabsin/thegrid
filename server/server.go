package main

import (
	"archive/zip"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var registeredApps []string

func registerApp(name, zipPath string) {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Fatalf("failed to open %s: %v", zipPath, err)
	}
	http.Handle(fmt.Sprintf("/%s/", name), http.StripPrefix(fmt.Sprintf("/%s/", name), http.FileServer(http.FS(zipReader))))
	registeredApps = append(registeredApps, name)
}

func main() {
	registerApp("solitaire", "solitaire/solitaire_pkg.zip")
	registerApp("example", "example/example_pkg.zip")
	registerApp("animdemo", "animdemo/animdemo_pkg.zip")

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, registeredApps)
	})

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}