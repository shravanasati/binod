package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w)
		return
	}

	if r.Method == "GET" {
		// files := []string{
		// 	"./web/templates/index.html",
		// 	"./web/templates/navbase.html",
		// 	"./web/templates/footer.html",
		// }
		tmpl := template.Must(template.New("404.html").Parse(indexTemplate))
		template.Must(tmpl.Parse(footerTemplate))
		template.Must(tmpl.Parse(navbaseTemplate))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func notFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	// files := []string{
	// 	"./web/templates/404.html",
	// 	"./web/templates/navbase.html",
	// 	"./web/templates/footer.html",
	// }
	tmpl := template.Must(template.New("404.html").Parse(notFoundTemplate))
	template.Must(tmpl.Parse(footerTemplate))
	template.Must(tmpl.Parse(navbaseTemplate))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
	}
}

func leaderBoardPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// all files
		// files := []string{
		// 	"./web/templates/leaderboard.html",
		// 	"./web/templates/navbase.html",
		// 	"./web/templates/footer.html",
		// }

		leaderboardData := make(map[int]Player)
		for i, v := range getLeaderBoardData() {
			leaderboardData[i+1] = v
		}

		// templates making and parsing
		tmpl := template.Must(template.New("leaderboard.html").Parse(leaderboardTemplate))
		template.Must(tmpl.Parse(footerTemplate))
		template.Must(tmpl.Parse(navbaseTemplate))

		// execute template
		if err := tmpl.Execute(w, leaderboardData); err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.New("login.html").Parse(signUpTemplate))
	template.Must(tmpl.Parse(footerTemplate))
	template.Must(tmpl.Parse(navbaseTemplate))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
	}
}

func joinPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// check for http method
	if end := checkForInvalidMethod("POST", r, w); end {
		return
	}

	// decoding json data
	player, end := decodeJSONToPlayer(r.Body, w)
	if end {
		return
	}

	// check for empty data
	if checkForEmpty(player, w) {
		return
	}

	// adding a player to the database
	success, message := addPlayer(player)
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

	// making response
	var resp *PlayerResponse
	if success {
		resp = makePlayerResponse(success, message, map[string]interface{}{
			"username": player.Username,
			"binods":   player.Binods,
		})
	} else {
		resp = makePlayerResponse(success, message, map[string]interface{}{})
	}

	w.Write(jsonifyResponse(resp))
}

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// check for http method
	if end := checkForInvalidMethod("GET", r, w); end {
		return
	}

	username := r.URL.Query().Get("username")
	if strip(username) == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := makePlayerResponse(false, "Username is required.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return
	}

	// getting player data
	player, err := getPlayer(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := makePlayerResponse(false, err.Error(), map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return
	}

	// making response
	resp := makePlayerResponse(true, fmt.Sprintf("A player with username '%s' found.", username), map[string]interface{}{
		"username": player.Username,
		"binods":   player.Binods,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(jsonifyResponse(resp))
}

func removePlayerHandler(w http.ResponseWriter, r *http.Request) {
	// check for valid http method
	if end := checkForInvalidMethod("DELETE", r, w); end {
		return
	}

	// decoding json data
	player, end := decodeJSONToPlayer(r.Body, w)
	if end {
		return
	}

	// check for empty data
	if checkForEmpty(player, w) {
		return
	}

	// delete a player from the database
	if ok := deletePlayer(player); !ok {
		w.WriteHeader(http.StatusForbidden)
		resp := makePlayerResponse(false, "Username or password is incorrect.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return
	}

	// making response
	w.WriteHeader(http.StatusOK)
	resp := makePlayerResponse(true, "Successfully deleted player.", map[string]interface{}{
		"username": player.Username,
		"binods":   player.Binods,
	})
	w.Write(jsonifyResponse(resp))
}

func updatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	// check for valid http method
	if end := checkForInvalidMethod("POST", r, w); end {
		return
	}

	// decoding json data
	player, end := decodeJSONToPlayer(r.Body, w)
	if end {
		return
	}

	// check for empty data
	if checkForEmpty(player, w) {
		return
	}

	// delete a player from the database
	if ok := updatePlayer(player); !ok {
		w.WriteHeader(http.StatusForbidden)
		resp := makePlayerResponse(false, "Username or password is incorrect.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return
	}

	// making response
	w.WriteHeader(http.StatusOK)
	resp := makePlayerResponse(true, "Successfully updated player.", map[string]interface{}{
		"username": player.Username,
		"binods":   player.Binods,
	})
	w.Write(jsonifyResponse(resp))
}

func leaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	if end := checkForInvalidMethod("GET", r, w); end {
		return
	}

	// getting leaderboard data
	leaderboard := getLeaderBoardData()

	data := make(map[int]interface{})
	for i, p := range leaderboard {
		data[i+1] = map[string]interface{}{
			"username": p.Username,
			"binods":   p.Binods,
		}
	}

	// making response
	resp := &LeaderboardResponse{
		Success: true,
		Data:    data,
		Message: "Successfully retrieved leaderboard.",
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonifyResponse(resp))
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

	postMessage(username, message)

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
	// website
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/leaderboardPage", leaderBoardPageHandler)
	http.HandleFunc("/registerPage", registerPageHandler)

	// player api
	http.HandleFunc("/player/join", joinPlayerHandler)
	http.HandleFunc("/player/get", getPlayerHandler)
	http.HandleFunc("/player/update", updatePlayerHandler)
	http.HandleFunc("/player/remove", removePlayerHandler)
	http.HandleFunc("/leaderboard", leaderBoardHandler)

	// message api
	http.HandleFunc("/message/post", postMessageHandler)
	http.HandleFunc("/message/get", getMessageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
