package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dre4success/tfl-underground-lines/utils"
)

type arrivals struct {
	NaptanId        string `json:"naptanId"`
	StationName     string `json:"stationName"`
	LineId          string `json:"lineId"`
	LineName        string `json:"lineName"`
	PlatformName    string `json:"platformName"`
	Towards         string `json:"towards"`
	DestinationName string `json:"destinationName"`
	CurrentLocation string `json:"currentLocation"`
	ExpectedArrival string `json:"expectedArrival"`
}

func (tl TflLines) StopArrivals(w http.ResponseWriter, r *http.Request) {
	naptanId := r.PathValue("naptanID")
	arrivalData := utils.TflData{
		URL: BASE_URL + "/StopPoint/" + naptanId + "/Arrivals",
	}
	data, err := arrivalData.FetchData()
	if err != nil {
		fmt.Printf("error fetching data: %s\n", err)
		return
	}
	var stopArrivals []arrivals
	err = json.Unmarshal(data, &stopArrivals)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	calCulateEstimatedArrival(&stopArrivals)

	jsonResponse, err := json.Marshal(stopArrivals)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func calCulateEstimatedArrival(arrive *[]arrivals) {
	for i := range *arrive {

		arrivalTime, err := time.Parse(time.RFC3339, (*arrive)[i].ExpectedArrival)
		if err != nil {
			log.Println("Error: invalid time format")
			continue
		}

		duration := time.Until(arrivalTime)

		if duration < 0 {
			(*arrive)[i].ExpectedArrival = "Already arrived"
			continue
		}

		if duration < time.Minute {
			(*arrive)[i].ExpectedArrival = "Less than a minute away"
			continue
		}

		minutes := int(duration.Minutes())
		hours := minutes / 60
		minutes = minutes % 60
		if hours > 0 {
			(*arrive)[i].ExpectedArrival = fmt.Sprintf("%d hour%s and %d minute%s away", hours, pluralize(hours), minutes, pluralize(minutes))
		} else {
			(*arrive)[i].ExpectedArrival = fmt.Sprintf("%d minute%s away", minutes, pluralize(minutes))
		}
	}
}

func pluralize(t int) string {
	if t == 1 {
		return ""
	}
	return "s"
}
