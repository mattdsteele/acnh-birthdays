package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mattdsteele/acnh-bdays"
)

func main() {
	http.HandleFunc("/", acnh.AcnhGopher)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		n := r.URL.Query().Get("name")
		fmt.Fprintf(w, "Hi %s!", n)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
