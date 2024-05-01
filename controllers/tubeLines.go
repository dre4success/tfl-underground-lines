package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dre4success/tfl-underground-lines/templates"
	"github.com/dre4success/tfl-underground-lines/utils"
)

type tube struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	LineStatuses []lineStatus `json:"lineStatuses"`
}

type lineStatus struct {
	StatusSeverity            int    `json:"statusSeverity"`
	StatusSeverityDescription string `json:"statusSeverityDescription"`
	Reason                    string `json:"reason"`
}

type TflLines struct {
	Lines templates.Template
	LineRoutes templates.Template
}

var lineColors = map[string]string{
	"bakerloo":         "#B36305",
	"central":          "#E32017",
	"circle":           "#FFD300",
	"district":         "#00782A",
	"hammersmith-city": "#F3A9BB",
	"jubilee":          "#A0A5A9",
	"metropolitan":     "#9B0056",
	"northern":         "#000000",
	"piccadilly":       "#003688",
	"victoria":         "#0098D4",
	"waterloo-city":    "#95CDBA",
}

const BASE_URL = "https://api.tfl.gov.uk"

func (tl TflLines) DisplayLines(w http.ResponseWriter, r *http.Request) {
	lineData := utils.TflData{
		URL: BASE_URL + "/line/mode/tube/status",
	}
	data, err := lineData.FetchData()
	if err != nil {
		fmt.Printf("error fetching data: %s\n", err)
		return
	}
	var lineTubes []tube
	if err := json.Unmarshal(data, &lineTubes); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	tplData := struct {
		Lines      []tube
		LineColors map[string]string
	}{
		Lines:      lineTubes,
		LineColors: lineColors,
	}

	if err := tl.Lines.Execute(w, tplData); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}
