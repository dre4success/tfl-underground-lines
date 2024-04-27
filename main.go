package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dre4success/tfl-underground-lines/controllers"
	"github.com/dre4success/tfl-underground-lines/templates"
)

func main() {

	lineC := controllers.TflLines{}
	lineC.Lines = templates.Must(templates.ParseFS("lines.gohtml"))

	fs := http.FileServer(http.Dir("static"))
	
	http.HandleFunc("/", lineC.DisplayLines)
	
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("starting server on :6060...")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
