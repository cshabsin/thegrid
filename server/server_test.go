package main

import (
	"archive/zip"
	"bytes"
	
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"testing/fstest"
)

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
	h := &appHandler{name: "testapp", zipReader: zipReader, firebaseConfig: nil, templates: nil}

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
	templates := fstest.MapFS{
		"layout.html.tpl": &fstest.MapFile{
			Data: []byte("{{template \"auth_ui\" .}}{{template \"body\" .}}"),
		},
		"auth_ui.html.tpl": &fstest.MapFile{
			Data: []byte("{{define \"auth_ui\"}}<div id=\"auth-container\">auth ui</div>{{end}}"),
		},
	}
	h := &appHandler{name: "testapp", zipReader: zipReader, firebaseConfig: nil, templates: templates}

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