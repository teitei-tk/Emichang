package weather

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient() (*Client, error) {
	r, err := url.ParseRequestURI(API_URL)
	if err != nil {
		return nil, err
	}

	logger := log.New(ioutil.Discard, "", log.LstdFlags)

	return &Client{
		URL:        r,
		HTTPClient: http.DefaultClient,
		Logger:     logger,
	}, nil
}
