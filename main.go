package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/melsincostan/brightsky-temp/metrics"
)

const (
	latitudeEnv  = "LATITUDE"
	longitudeEnv = "LONGITUDE"
)

func main() {
	rlat, ok := os.LookupEnv(latitudeEnv)
	if !ok {
		log.Fatalf("missing latitude (%s)", latitudeEnv)
	}

	rlon, ok := os.LookupEnv(longitudeEnv)
	if !ok {
		log.Fatalf("missing longitude (%s)", longitudeEnv)
	}

	lat, err := strconv.ParseFloat(rlat, 64)
	if err != nil {
		log.Fatalf("could not parse latitude: %s", err.Error())
	}

	lon, err := strconv.ParseFloat(rlon, 64)
	if err != nil {
		log.Fatalf("could not parse longitude: %s", err.Error())
	}

	log.Printf("starting for lat: %f, lon: %f", lat, lon)
	http.Handle("GET /metrics", metrics.Handler(lat, lon))

	http.ListenAndServe(":8080", nil)
}
