package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dre4success/tfl-underground-lines/utils"
)

type routes struct {
	LineId             string               `json:"lineId"`
	LineName           string               `json:"lineName"`
	LineStrings        []string             `json:"lineStrings"`
	Stations           []stations           `json:"stations"`
	StopPointSequences []stopPointSequences `json:"stopPointSequences"`
}

type stations struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Lat       float64    `json:"lat"`
	Lon       float64    `json:"lon"`
	StationId string `json:"stationId"`
}

type stopPointSequences struct {
	StopPoint []stopPoint `json:"stopPoint"`
}

type stopPoint struct {
	ParentId  string `json:"parentId"`
	StationId string `json:"stationId"`
}

func (tl TflLines) Routes(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("ID")
	lineData := utils.TflData{
		URL: BASE_URL + "/line/" + id + "/route/sequence/outbound",
	}
	data, err := lineData.FetchData()
	if err != nil {
		fmt.Printf("error fetching data: %s\n", err)
		return
	}

	var lineRoutes routes
	err = json.Unmarshal(data, &lineRoutes)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	tl.LineRoutes.Execute(w, lineRoutes)
}
