package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// from http://weather.livedoor.com/weather_hacks/webservice
type Response struct {
	PinpointLocations []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"pinpointLocations"`
	Link      string `json:"link"`
	Forecasts []struct {
		DateLabel   string `json:"dateLabel"`
		Telop       string `json:"telop"`
		Date        string `json:"date"`
		Temperature struct {
			Min interface{} `json:"min"`
			Max interface{} `json:"max"`
		} `json:"temperature"`
		Image struct {
			Width  int    `json:"width"`
			URL    string `json:"url"`
			Title  string `json:"title"`
			Height int    `json:"height"`
		} `json:"image"`
	} `json:"forecasts"`
	Location struct {
		City       string `json:"city"`
		Area       string `json:"area"`
		Prefecture string `json:"prefecture"`
	} `json:"location"`
	PublicTime string `json:"publicTime"`
	Copyright  struct {
		Provider []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"provider"`
		Link  string `json:"link"`
		Title string `json:"title"`
		Image struct {
			Width  int    `json:"width"`
			Link   string `json:"link"`
			URL    string `json:"url"`
			Title  string `json:"title"`
			Height int    `json:"height"`
		} `json:"image"`
	} `json:"copyright"`
	Title       string `json:"title"`
	Description struct {
		Text       string `json:"text"`
		PublicTime string `json:"publicTime"`
	} `json:"description"`
}

func Request(client *Client) (*Response, error) {
	res := &Response{}

	req, err := http.NewRequest(http.MethodGet, client.URL.String(), nil)
	if err != nil {
		return res, err
	}

	result, err := client.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, res)
	if err != nil {
		return res, err
	}

	return res, nil
}
