package main

import (
	"archive/zip"
	"fmt"
	"log"
	"net/http"
)

func registerApp(name, zipPath string) {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Fatalf("failed to open %s: %v", zipPath, err)
	}
	http.Handle(fmt.Sprintf("/%s/", name), http.StripPrefix(fmt.Sprintf("/%s/", name), http.FileServer(http.FS(zipReader))))
}

func main() {
	registerApp("solitaire", "solitaire/solitaire_pkg.zip")
	registerApp("example", "example/example_pkg.zip")

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}