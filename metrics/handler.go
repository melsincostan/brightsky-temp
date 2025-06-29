package metrics

import (
	"log"
	"net/http"

	"github.com/melsincostan/brightsky-temp/brightsky"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler(lat, lon float64) http.HandlerFunc {
	ph := promhttp.Handler()
	return func(w http.ResponseWriter, r *http.Request) {
		curr_val, err := brightsky.Data(lat, lon)
		if err != nil {
			log.Printf("Error retrieving data: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		Temperature.Set(force(curr_val.Weather.Temperature))
		RelativeHumidity.Set(float64(force(curr_val.Weather.RelativeHumidity)))
		CloudCover.Set(force(curr_val.Weather.CloudCover))
		PressureMSL.Set(force(curr_val.Weather.PressureMSL))
		WindSpeed.Set(force(curr_val.Weather.WindSpeed10))

		ph.ServeHTTP(w, r)
	}
}

func force[T any](val *T) (res T) {
	if val == nil {
		return
	}

	return *val
}
