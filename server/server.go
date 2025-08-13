package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cshabsin/thegrid/apps/explorers/data"
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

var config struct {
	Firebase struct {
		APIKey            string `json:"apiKey"`
		AuthDomain        string `json:"authDomain"`
		ProjectID         string `json:"projectId"`
		StorageBucket     string `json:"storageBucket"`
		MessagingSenderID string `json:"messagingSenderId"`
		AppID             string `json:"appId"`
	} `json:"firebase"`
}

func main() {
	configFile, err := os.Open("server/config.json")
	if err != nil {
		log.Fatal("failed to open config file: ", err)
	}
	defer configFile.Close()
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatal("failed to decode config file: ", err)
	}

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
					<script>
						var firebaseConfig = {{.FirebaseConfig}};
					</script>
				</head>
				<body>
					<h1>Available Services</h1>
					<ul>
						{{range .RegisteredApps}}
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
		firebaseConfigJSON, err := json.Marshal(config.Firebase)
		if err != nil {
			log.Printf("json marshal error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			RegisteredApps []string
			FirebaseConfig template.JS
		}{
			RegisteredApps: registeredApps,
			FirebaseConfig: template.JS(firebaseConfigJSON),
		}
		if err := t.Execute(w, data); err != nil {
			log.Printf("template execute error: %v", err)
		}
	})

	http.HandleFunc("/firebase/auth/bundle.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "firebase/auth/bundle.js")
	})

	http.HandleFunc("/explorers/data", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(data.ExplorersMapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
