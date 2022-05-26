// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/cshabsin/thegrid/example/server/data"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(data.ExplorersMapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	err := http.ListenAndServe(*listen, nil)
	log.Fatalln(err)
}
