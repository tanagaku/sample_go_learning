package main

import (
	"encoding/json"
	"net/http"
)

func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
