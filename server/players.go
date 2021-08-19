package main

import (
	"fmt"
	"sort"
	"sync"
)

type Player struct {
	Username string
	password string
	Binods int
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
		if player.Username == username {
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
	binodList := make([]int, len(playerdb.players))
	for i, player := range playerdb.players {
		binodList[i] = player.Binods
	}

	sort.Ints(binodList)

	data := "Binod Leaderboard: \n"
	for _, v := range binodList {
		data += fmt.Sprintf("%v: %v binods \n", playerdb.players[v].Username, playerdb.players[v].Binods)
	}

	return data
}

func getLeaderBoardData() []Player {
	playerdb.Lock()
	defer playerdb.Unlock()

	playersSorted := make([]Player, len(playerdb.players))
	
	for i, player := range playerdb.players {
		playersSorted[i] = player
	}

	sort.Slice(playersSorted, func(i, j int) bool {
		return playersSorted[i].Binods > playersSorted[j].Binods
	})

	return playersSorted
}

func updatePlayer(username, password string, binodCount int) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.Username == username {
			playerdb.players[i].Binods = binodCount
			return true
		}
	}

	return false
}

func deletePlayer(username, password string) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.Username == username && player.password == password {
			playerdb.players = append(playerdb.players[:i], playerdb.players[i+1:]...)
			return true
		}
	}

	return false
}