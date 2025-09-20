package shawapi

import "net/http"

// These structs will define how the json item returned from shaw will look like to us,
// After we have unmarshalled it from shaw endpoint.
type Movie struct {
	Title		string `json:"primaryTitle"`
	PosterUrl 	string `json:"posterUrl"`
	Duration	string `json:"duration"`
	ShowTimes	MovieTiming
}

type MovieTiming struct {
	ShowDate	string `json:"displayDate"`
	ShowTime	string `json:"displayTime"`
	SeatingStatus string `json:"seatingStatus"`
}

const shawBaseUrl = "https://shaw.sg"

// This will help us call the shaw API.
type ShawClient struct {
	BaseUrl string
	HTTPClient *http.Client
}

