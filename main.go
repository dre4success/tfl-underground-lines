package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dre4success/tfl-underground-lines/controllers"
	"github.com/dre4success/tfl-underground-lines/templates"
)

func main() {

	b := "base.gohtml"

	lineC := controllers.TflLines{}
	lineC.Lines = templates.Must(templates.ParseFS(b, "lines.gohtml"))
	lineC.LineRoutes = templates.Must(templates.ParseFS(b, "lineRoutes.gohtml"))

	fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/", lineC.DisplayLines)
	http.HandleFunc("/line/{ID}", lineC.Routes)

	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("starting server on :6060...")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
