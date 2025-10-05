package shawapi

import (
	"net/http"
	"time"
)

// These structs will define how the json item returned from shaw will look like to us,
// After we have unmarshalled it from shaw endpoint.
type Movie struct {
	Title		string `json:"primaryTitle"`
	PosterUrl 	string `json:"posterUrl"`
	Duration	string `json:"duration"`
	ShowTimes	[]MovieTiming `json:"showTimes"`
}

type MovieTiming struct {
	DisplayDate	string `json:"displayDate"`
	DisplayTime	string `json:"displayTime"`
	SeatingStatus string `json:"seatingStatus"`
}

const shawBaseUrl = "https://shaw.sg"

// This will help us call the shaw API.
type ShawClient struct {
	BaseUrl string
	HTTPClient *http.Client
}

// go idiom: for this constructor, we choose to return ShawClient as an address because we are going to pass
// this struct around, and not creating copies on the fly. Thus to be memory savvy and to retain info in the client as we pass it down the function calls, we refer to the same address throughout the client call being made to shaw.
func NewClient() *ShawClient {
	
	return &ShawClient{
		BaseUrl: shawBaseUrl,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}