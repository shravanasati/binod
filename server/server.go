package main

import (
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

func joinHandler(w http.ResponseWriter, r *http.Request) {
	// check for http method
	if end := checkForInvalidMethod("POST", r, w); end {
		return
	}

	player, end := decodeJSONToPlayer(r.Body, w)
	if end {
		return
	}

	if checkForEmpty(player, w) {
		return
	}

	success, message := addPlayer(player)
	if !success {
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	var resp *PlayerResponse
	if success {
		resp = makeResponse(success, message, map[string]interface{}{
			"username": player.Username,
			"binods":   player.Binods,
		})
	} else {
		resp = makeResponse(success, message, map[string]interface{}{})
	}

	w.Write(jsonifyResponse(resp))
}

func removeProfileHandler(w http.ResponseWriter, r *http.Request) {
	if end := checkForInvalidMethod("DELETE", r, w); end {
		return
	}

	player, end := decodeJSONToPlayer(r.Body, w)
	if end {
		return
	}

	if checkForEmpty(player, w) {
		return
	}

	if ok := deletePlayer(player); !ok {
		w.WriteHeader(http.StatusForbidden)
		resp := makeResponse(false, "Username or password is incorrect.", map[string]interface{}{})
		w.Write(jsonifyResponse(resp))
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := makeResponse(true, "Successfully deleted player.", map[string]interface{}{
		"username": player.Username,
		"binods":   player.Binods,
	})
	w.Write(jsonifyResponse(resp))
}

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	binod := r.URL.Query().Get("binod")



	if ok := updatePlayer(&player); !ok {
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

	// player api
	http.HandleFunc("/join", joinHandler)
	http.HandleFunc("/update", updateProfileHandler)
	http.HandleFunc("/remove", removeProfileHandler)
	http.HandleFunc("/leaderboard", leaderBoardHandler)

	// message api
	http.HandleFunc("/postmessage", postMessageHandler)
	http.HandleFunc("/getmessage", getMessageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
