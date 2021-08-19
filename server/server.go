package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		jsonStr, e := json.Marshal(map[string]string{
			"message": "Welcome",
		})
		if e != nil {
			log.Fatal(e)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	// check for http method
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// retrieve request queries
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	binod := r.URL.Query().Get("binod")

	// check for empty queries
	if username == "" || password == "" || binod == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
		return
	}

	// make sure binod is a number
	binodCount, e := strconv.Atoi(binod)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
		return
	}
	
	// making a new player
	if !newPlayer(username, password, binodCount) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Username already exists"))
		return
	}

	w.Write([]byte("Welcome " + username))
}

func removeProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
		return
	}

	if ok := deletePlayer(username, password); !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Username not found or invalid credentials"))
		return
	}

	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Invalid credentials"))
}

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	binod := r.URL.Query().Get("binod")

	if username == "" || password == "" || binod == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
		return
	}

	binodCount, e := strconv.Atoi(binod)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
		return
	}
	
	if ok := updatePlayer(username, password, binodCount); !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid credentials"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated profile " + username))
}

func leaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data := getLeaderBoard()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	message := r.URL.Query().Get("message")

	if message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
		return
	}

	postMessage(username,message)

	w.Write([]byte("Message posted"))
}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	message := getMessage()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/join", joinHandler)
	http.HandleFunc("/update", updateProfileHandler)
	http.HandleFunc("/remove", removeProfileHandler)
	http.HandleFunc("/leaderboard", leaderBoardHandler)

	http.HandleFunc("/postmessage", postMessageHandler)
	http.HandleFunc("/getmessage", getMessageHandler)

	log.Fatal(http.ListenAndServe(":6969", nil))
}