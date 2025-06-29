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
	WindDirection = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_wind_direction_10",
		Help: "Wind direction over the last 10 minutes",
	})
	WindGustSpeed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_wind_gust_speed_10",
		Help: "Average wind gust speed over the last 10 minutes",
	})
	WindGustDirection = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_wind_gust_direction_10",
		Help: "Wind gust direction over the last 10 minutes",
	})
	Precipitation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_precipitation_10",
		Help: "precipitation over the last 10 minutes",
	})
	DewPoint = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_dew_point",
		Help: "Dew Point",
	})
	SolarIrradiation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_solar_irradiation_10",
		Help: "Solar Irradiation over the last 10 minues",
	})
	SunshineTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bs_curr_sunshine_30",
		Help: "Sunshine over the last 30 minutes",
	})
)
