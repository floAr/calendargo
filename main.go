package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement me!
		http.Error(w, fmt.Errorf("not implemented").Error(), http.StatusInternalServerError)
	})

}
