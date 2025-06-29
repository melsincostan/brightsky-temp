package brightsky

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	urlFormat = "/current_weather?lat=%f&lon=%f"
)

var host = "https://api.brightsky.dev"

var client = http.Client{
	Timeout: 10 * time.Second,
}

var (
	ErrAPI = errors.New("non-200 return code from the API")
)

func Data(lat, lon float64) (data *Response, err error) {
	endpoint := fmt.Sprintf(urlFormat, lat, lon)
	url := fmt.Sprintf("%s%s", host, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w (%s)", ErrAPI, res.Status)
	}

	data = new(Response)
	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}
	return
}

func SetHost(host string) {
	host = strings.TrimSuffix(host, "/")
}

func SetClient(newClient http.Client) {
	client = newClient
}
