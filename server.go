package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	serveError := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if serveError != nil {
		log.Print(serveError)
		log.Fatal("Oh dear, the server exploded!")
	}
}
