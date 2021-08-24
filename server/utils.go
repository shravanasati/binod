package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type PlayerResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func strip(data string) string {
	return strings.TrimSpace(data)
}

func makeResponse(success bool, message string, data map[string]interface{}) *PlayerResponse {
	response := &PlayerResponse{
		Success: success,
		Message: message,
		Data:    data,
	}

	return response
}

func jsonifyResponse(resp *PlayerResponse) []byte {
	json, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error marshalling response: %s", err)
	}

	return json
}

// check for empty username and password and return true if the response has been written.
func checkForEmpty(p *Player, w http.ResponseWriter) bool {
	if strip(p.Username) == "" || strip(p.Password) == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := makeResponse(false, "Username or password cannot be empty.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return true
	}
	return false
}

func decodeJSONToPlayer(content io.Reader, w http.ResponseWriter) (*Player, bool) {
	decoder := json.NewDecoder(content)
	var player Player
	if err := decoder.Decode(&player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := makeResponse(false, "Invalid JSON object.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return nil, true
	}
	return &player, false
}

func checkForInvalidMethod(acceptedMethod string, r *http.Request, w http.ResponseWriter) bool {
	if r.Method != acceptedMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)

		resp := makeResponse(false, fmt.Sprintf("Method '%s' not allowed. The only valid HTTP method for this endpoint is '%s'", r.Method, acceptedMethod), map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return true
	}
	return false
}