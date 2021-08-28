package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// PlayerResponse is a struct that holds the response data for the player api handlers.
type PlayerResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LeaderboardResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    map[int]interface{} `json:"data"`
}

// strip removes all whitespace from the given string.
func strip(data string) string {
	return strings.TrimSpace(data)
}

func makePlayerResponse(success bool, message string, data map[string]interface{}) *PlayerResponse {
	response := &PlayerResponse{
		Success: success,
		Message: message,
		Data:    data,
	}

	return response
}

// jsonifyResponse takes a PlayerResponse object and returns a JSON string.
func jsonifyResponse(resp interface{}) []byte {
	jsonStr, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error marshalling response: %s", err)
	}

	return jsonStr
}

// check for empty username and password and return true if the response has been written.
func checkForEmpty(p *Player, w http.ResponseWriter) bool {
	if strip(p.Username) == "" || strip(p.Password) == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := makePlayerResponse(false, "Username or password cannot be empty.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return true
	}
	return false
}

// decodeJSONToPlayer decodes the given JSON into a Player struct and if the conversion
// failed writes a proper response and returns the boolean value if to continue function
// execution or not.
func decodeJSONToPlayer(content io.Reader, w http.ResponseWriter) (*Player, bool) {
	decoder := json.NewDecoder(content)
	var player Player
	if err := decoder.Decode(&player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := makePlayerResponse(false, "Invalid JSON object.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return nil, true
	}
	return &player, false
}

// checkForInvalidMethod takes the acceptedMethod as a string and checks if the method is
// valid. If its not valid, it writes the proper response. Returns a boolean value for
// whether to end the function body or not (since the response is already written).
func checkForInvalidMethod(acceptedMethod string, r *http.Request, w http.ResponseWriter) bool {
	if r.Method != acceptedMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)

		resp := makePlayerResponse(false, fmt.Sprintf("Method '%s' not allowed. The only valid HTTP method for this endpoint is '%s'", r.Method, acceptedMethod), map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return true
	}
	return false
}
