package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi there!I am go webserver")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
