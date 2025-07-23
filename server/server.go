package main

import (
	"archive/zip"
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var dataDir = flag.String("data_dir", ".", "directory containing the zip files")

func main() {
	flag.Parse()

	solitaireZipPath := filepath.Join(*dataDir, "solitaire.zip")
	exampleZipPath := filepath.Join(*dataDir, "example.zip")

	solitaireZipReader, err := zip.OpenReader(solitaireZipPath)
	if err != nil {
		log.Fatalf("failed to open solitaire.zip: %v", err)
	}

	exampleZipReader, err := zip.OpenReader(exampleZipPath)
	if err != nil {
		log.Fatalf("failed to open example.zip: %v", err)
	}

	http.Handle("/solitaire/", http.StripPrefix("/solitaire/", http.FileServer(http.FS(solitaireZipReader))))
	http.Handle("/example/", http.StripPrefix("/example/", http.FileServer(http.FS(exampleZipReader))))

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
