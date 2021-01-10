package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Result struct {
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
	Status int         `json:"status"`
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	log.Println(payload)
	result := Result{}
	result.Data = payload
	result.Status = status
	response, err := json.Marshal(result)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	log.Println(message)
	result := Result{}
	result.Status = code
	result.Error = message
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(response))

}
