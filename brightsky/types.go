package brightsky

type Response struct {
	Weather Weather  `json:"weather"`
	Sources []Source `json:"sources"`
}

type Weather struct {
	Timestamp           string      `json:"timestamp"` // ISO 8601
	SourceID            int         `json:"source_id"`
	CloudCover          *float64    `json:"cloud_cover,omitempty"`
	Condition           *string     `json:"condition,omitempty"`
	DewPoint            *float64    `json:"dew_point,omitempty"`
	Icon                *string     `json:"icon,omitempty"`
	PressureMSL         *float64    `json:"pressure_msl,omitempty"`
	RelativeHumidity    *int        `json:"relative_humidity,omitempty"`
	Temperature         *float64    `json:"temperature,omitempty"`
	Visibility          *int        `json:"visibility,omitempty"`
	FallbackSourceIds   interface{} `json:"fallback_source_ids"`
	Precipitation10     *float64    `json:"precipitation_10,omitempty"`
	Precipitation30     *float64    `json:"precipitation_30,omitempty"`
	Precipitation60     *float64    `json:"precipitation_60,omitempty"`
	Solar10             *float64    `json:"solar_10,omitempty"`
	Solar30             *float64    `json:"solar_30,omitempty"`
	Solar60             *float64    `json:"solar_60,omitempty"`
	Sunshine30          *float64    `json:"sunshine_30,omitempty"`
	Sunshine60          *float64    `json:"sunshine_60,omitempty"`
	WindDirection10     *int        `json:"wind_direction_10,omitempty"`
	WindDirection30     *int        `json:"wind_direction_30,omitempty"`
	WindDirection60     *int        `json:"wind_direction_60,omitempty"`
	WindSpeed10         *float64    `json:"wind_speed_10,omitempty"`
	WindSpeed30         *float64    `json:"wind_speed_30,omitempty"`
	WindSpeed60         *float64    `json:"wind_speed_60,omitempty"`
	WindGustDirection10 *int        `json:"wind_gust_direction_10,omitempty"`
	WindGustDirection30 *int        `json:"wind_gust_direction_30,omitempty"`
	WindGustDirection60 *int        `json:"wind_gust_direction_60,omitempty"`
	WindGustSpeed10     *float64    `json:"wind_gust_speed_10,omitempty"`
	WindGustSpeed30     *float64    `json:"wind_gust_speed_30,omitempty"`
	WindGustSpeed60     *float64    `json:"wind_gust_speed_60,omitempty"`
}

type Source struct {
	ID              int     `json:"id"`
	DWDStationID    *string `json:"dwd_station_id,omitempty"`
	WMOStationID    *string `json:"wmo_station_id,omitempty"`
	StationName     *string `json:"station_name"`
	ObservationType string  `json:"observation_type"`
	FirstRecord     string  `json:"first_record"`
	LastRecord      string  `json:"last_record"`
	Latitude        float64 `json:"lat"`
	Longitude       float64 `json:"long"`
	Height          float64 `json:"height"`
	Distance        float64 `json:"distance"`
}
