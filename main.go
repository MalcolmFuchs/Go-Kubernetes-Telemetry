package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TelematikData struct {
	VehicleID string  `json:"vehicle_id"`
	Speed     float64 `json:"speed"`
	Location  string  `json:"location"`
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	var data TelematikData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	log.Printf("Received data: %+v\n", data)
	fmt.Fprintf(w, "Data received\n")
}

func main() {
	http.HandleFunc("/telematik", dataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
