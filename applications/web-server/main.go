package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "Port to listen on")
	folder := flag.String("folder", ".", "Folder to serve")
	flag.Parse()

	log.Println("Starting...")

	http.Handle("/", http.FileServer(http.Dir(*folder)))

	log.Printf("Listening on http://localhost:%v \n", *port)
	log.Printf("Serving folder %v \n", *folder)

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
