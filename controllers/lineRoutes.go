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
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	StationId string  `json:"stationId"`
}

type stopPointSequences struct {
	StopPoint []stopPoint `json:"stopPoint"`
}

type stopPoint struct {
	TopMostParentId string `json:"topMostParentId"`
	StationId       string `json:"stationId"`
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
	mapStationIDs(&lineRoutes)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	tl.LineRoutes.Execute(w, lineRoutes)
}

// geting the stationIds as some stations have IDs that are different from stationIds.
// The stationIds become useful when we want to get stop arrivals, as they can be passed
// to a stop arrival endpoint as the naptanId.
func mapStationIDs(routesData *routes) {
	stationIDMap := make(map[string]string)
	for _, stSeq := range routesData.StopPointSequences {
		for _, sp := range stSeq.StopPoint {
			stationIDMap[sp.TopMostParentId] = sp.StationId
		}
	}

	// map station Ids to the top level stations
	for i, station := range routesData.Stations {
		stationID, ok := stationIDMap[station.Id]
		if ok {
			routesData.Stations[i].StationId = stationID
		}
	}
}
