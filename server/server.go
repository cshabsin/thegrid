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
			// Look for index.html.tpl
			indexTplFile, err := zipReader.Open("index.html.tpl")
			if err == nil {
				defer indexTplFile.Close()
				indexTplContent, err := io.ReadAll(indexTplFile)
				if err != nil {
					http.Error(w, "failed to read template", http.StatusInternalServerError)
					return
				}

				authTplFile, err := zipReader.Open("auth.html.tpl")
				if err != nil {
					http.Error(w, "failed to open auth template", http.StatusInternalServerError)
					return
				}
								defer authTplFile.Close()
				authTplContent, err := io.ReadAll(authTplFile)
				if err != nil {
					http.Error(w, "failed to read auth template", http.StatusInternalServerError)
					return
				}

				t, err := template.New("index").Parse(string(authTplContent))
				if err != nil {
					http.Error(w, "failed to parse auth template", http.StatusInternalServerError)
					return
				}

				t, err = t.Parse(string(indexTplContent))
				if err != nil {
					http.Error(w, "failed to parse index template", http.StatusInternalServerError)
					return
				}

				data := struct {
					FirebaseConfig any
				}{
					FirebaseConfig: config.Firebase,
				}
				if err := t.Execute(w, data); err != nil {
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
		data := struct {
			RegisteredApps []string
			FirebaseConfig any
		}{
			RegisteredApps: registeredApps,
			FirebaseConfig: config.Firebase,
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
