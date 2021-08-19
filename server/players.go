package main

import (
	"fmt"
	"sync"
)

type Player struct {
	username string
	password string
	binods int
}

type playerDB struct {
	players []Player
	sync.Mutex
}

var playerdb playerDB

// newPlayer adds a new player to the database if the username doesnt exist and returns a 
// boolean value of true if the username is unique.
func newPlayer(username, password string, binodCount int) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for _, player := range playerdb.players {
		if player.username == username {
			return false
		}
	}

	playerdb.players = append(playerdb.players, Player{username, password, binodCount})

	return true
}

// getLeaderBoard returns the player leaderboard.
func getLeaderBoard() string {
	playerdb.Lock()
	defer playerdb.Unlock()

	data := "Binod Leaderboard: \n"
	for _, player := range playerdb.players {
		data += fmt.Sprintf("%v: %v binods \n", player.username, player.binods)
	}

	return data
}

func getLeaderBoardData() []Player {
	playerdb.Lock()
	defer playerdb.Unlock()

	return playerdb.players
}

func updatePlayer(username, password string, binodCount int) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.username == username {
			playerdb.players[i].binods = binodCount
			return true
		}
	}

	return false
}

func deletePlayer(username, password string) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.username == username && player.password == password {
			playerdb.players = append(playerdb.players[:i], playerdb.players[i+1:]...)
			return true
		}
	}

	return false
}