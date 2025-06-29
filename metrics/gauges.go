package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "curr_temperature",
		Help: "Current Temperature",
	})
	RelativeHumidity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "curr_relative_humidity",
		Help: "Current Relative Humidity",
	})
	CloudCover = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_cloud_cover",
		Help: "Current Cloud Cover",
	})
	PressureMSL = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_pressure_msl",
		Help: "Current MSL Pressure",
	})
	WindSpeed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_wind_speed_10",
		Help: "Average wind speed over the last 10 minutes",
	})
)
