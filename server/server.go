package main

import (
	"archive/zip"
	"log"
	"net/http"
)

func main() {
	solitaireZipReader, err := zip.OpenReader("solitaire.zip")
	if err != nil {
		log.Fatalf("failed to open solitaire.zip: %v", err)
	}
	// defer solitaireZipReader.Close() // This would close the file before the server can use it.

	exampleZipReader, err := zip.OpenReader("example.zip")
	if err != nil {
		log.Fatalf("failed to open example.zip: %v", err)
	}
	// defer exampleZipReader.Close() // This would close the file before the server can use it.

	http.Handle("/solitaire/", http.StripPrefix("/solitaire/", http.FileServer(http.FS(solitaireZipReader))))
	http.Handle("/example/", http.StripPrefix("/example/", http.FileServer(http.FS(exampleZipReader))))

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}