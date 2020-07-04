package nepodate

import (
	"context"
	"encoding/json"
	"net/http"
)

type getRequest struct {}

type getResponse struct {
	Date string `json:"date"`
	Err string `json:"err,omitempty"`
}

type statusRequest struct {}

type statusResponse struct {
	Status string `json:"status"`
}

type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Valid bool `json:"valid"`
	Err string `json:"err,omitempty"`
}

// write decoders for incoming requests
// an empty interface (interface{}) is just an interface that is satisfied by all data types
func decodeGetRequest(ctx context.Context, r *http.Request)(interface{}, error)  {
	var req getRequest
	return req,nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request)(interface{},error)  {
	var req statusRequest
	return req, nil
}

func decodeValidateRequest(ctx context.Context, r *http.Request)(interface{},error)  {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req,nil
}

// encoder for response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)

}