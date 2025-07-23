package main

import (
	"archive/zip"
	"log"
	"net/http"
)

func main() {
	solitaireZipReader, err := zip.OpenReader("solitaire/solitaire_pkg.zip")
	if err != nil {
		log.Fatalf("failed to open solitaire.zip: %v", err)
	}

	exampleZipReader, err := zip.OpenReader("example/example_pkg.zip")
	if err != nil {
		log.Fatalf("failed to open example.zip: %v", err)
	}

	http.Handle("/solitaire/", http.StripPrefix("/solitaire/", http.FileServer(http.FS(solitaireZipReader))))
	http.Handle("/example/", http.StripPrefix("/example/", http.FileServer(http.FS(exampleZipReader))))

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}