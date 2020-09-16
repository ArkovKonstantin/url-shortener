package application

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type (
	ShortenRequest struct {
		URL string `json:"url"`
	}

	ShortenResponse struct {
		ShortURL string `json:"short_url"`
	}

	CreateRequest struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	}
)

func decodeShortenReq(r *http.Request) (*ShortenRequest, error) {
	var req ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	_, err = url.ParseRequestURI(req.URL)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func decodeCreateReq(r *http.Request) (*CreateRequest, error) {
	var req CreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	_, err = url.ParseRequestURI(req.URL)
	if err != nil {
		return nil, err
	}

	if req.Name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &req, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
