package server

import (
	"encoding/json"
	"net/http"
)

type SimpleHttpResponse struct {
	Success bool          `json:"success"`
	Data   DataMessage `json:"data"`
}

type DataMessage struct {
	Message	string	`json:"message"`
}

func SendHttpResp(w http.ResponseWriter, messageString string, statusCode int) {
	var requestStatus = true;
	if statusCode != 200 {
		requestStatus = false
	}

	jsonResponse := SimpleHttpResponse{
		Success: requestStatus,
		Data: DataMessage{
			Message: messageString,
		},
	}

	response, _ := json.Marshal(jsonResponse)

	w.WriteHeader(statusCode)
	w.Write(response)

	return
}