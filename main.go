package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

var (
	templateFile = template.Must(template.ParseFiles("static/index.html"))
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	templateFile.ExecuteTemplate(w, "index.html", nil)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad Requeat", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := path.Base(fileHeader.Filename)
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/?Succes=true", http.StatusSeeOther)
}
func main() {
	http.HandleFunc("/", uploadFile)

	log.Println("Server started")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
