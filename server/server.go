package main

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cshabsin/thegrid/apps/explorers/data"
	"github.com/cshabsin/thegrid/secretmanager"
)

var registeredApps []string
var zipReaders = make(map[string]*zip.ReadCloser)

func registerApp(name, zipPath string) {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Printf("failed to open %s: %v. Skipping.", zipPath, err)
		return
	}
	zipReaders[name] = zipReader

	fileServer := http.FileServer(http.FS(&zipReader.Reader))
	http.HandleFunc(fmt.Sprintf("/%s/", name), func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || strings.HasSuffix(r.URL.Path, "/index.html") {
			// Look for body.html.tpl
			bodyTplFile, err := zipReader.Open("body.html.tpl")
			if err == nil {
				defer bodyTplFile.Close()
				bodyTplContent, err := io.ReadAll(bodyTplFile)
				if err != nil {
					http.Error(w, "failed to read template", http.StatusInternalServerError)
					return
				}

				t, err := template.ParseFiles("server/templates/layout.html.tpl", "firebase/authui/auth_ui.html.tpl")
				if err != nil {
					http.Error(w, "failed to parse layout templates", http.StatusInternalServerError)
					return
				}

				t, err = t.New("body").Parse(string(bodyTplContent))
				if err != nil {
					http.Error(w, "failed to parse body template", http.StatusInternalServerError)
					return
				}

				data := struct {
					Title          string
					FirebaseConfig any
				}{
					Title:          name,
					FirebaseConfig: config.Firebase,
				}
				if err := t.ExecuteTemplate(w, "layout", data); err != nil {
					log.Printf("template execute error: %v", err)
				}
				return
			}
		}

		// Fallback to serving files from the zip archive
		http.StripPrefix(fmt.Sprintf("/%s/", name), fileServer).ServeHTTP(w, r)
	})

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
	ctx := context.Background()
	firebaseConfigJSON, err := secretmanager.GetSecret(ctx, "shabsin-thegrid", "firebase-config")
	if err != nil {
		log.Fatalf("failed to get firebase config from secret manager: %v", err)
	}
	if err := json.Unmarshal([]byte(firebaseConfigJSON), &config.Firebase); err != nil {
		log.Fatalf("failed to unmarshal firebase config: %v", err)
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
		data := struct {
			RegisteredApps []string
		}{
			RegisteredApps: registeredApps,
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