package main

import (
	"archive/zip"
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const layoutTpl = `
<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/firebase/authui/auth.css">
    <script>
        var firebaseConfig = {{.FirebaseConfig}};
    </script>
    {{template "head" .}}
</head>
<body>
    {{template "auth_ui" .}}
    <div id="content">
        {{template "body" .}}
    </div>
</body>
</html>
`

const authUiTpl = `
{{define "auth_ui"}}
<div id="auth-container">
    <div id="logged-out-view">
        <button id="login-button">Login</button>
    </div>
    <div id="logged-in-view" style="display: none;">
        <span id="user-name"></span>
        <button id="logout-button">Logout</button>
    </div>
</div>
{{end}}
`

func TestAppHandler_Static(t *testing.T) {
	// Create a test zip archive with a static index.html
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	fw, err := zw.Create("index.html")
	if err != nil {
		t.Fatal(err)
	}
	fw.Write([]byte("static index.html"))
	zw.Close()

	// Create a temporary file for the zip archive
	tmpfile, err := os.CreateTemp("", "test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(buf.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a test app handler
	zipReader, err := zip.OpenReader(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	h := &appHandler{name: "testapp", zipReader: zipReader, firebaseConfig: nil}

	req, err := http.NewRequest("GET", "/testapp/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := "static index.html"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAppHandler_Templated(t *testing.T) {
	// Create a test zip archive with an index.html.tpl
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	fw, err := zw.Create("index.html.tpl")
	if err != nil {
		t.Fatal(err)
	}
	fw.Write([]byte("{{define \"body\"}}template body{{end}}"))
	zw.Close()

	// Create a temporary file for the zip archive
	tmpfile, err := os.CreateTemp("", "test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(buf.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a test app handler
	zipReader, err := zip.OpenReader(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	h := &appHandler{name: "testapp", zipReader: zipReader, firebaseConfig: nil}

	// The test needs to parse the templates itself.
	tpl, err := template.New("layout").Parse(layoutTpl)
	if err != nil {
		t.Fatal(err)
	}
	tpl, err = tpl.Parse(authUiTpl)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/testapp/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := "template body"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}