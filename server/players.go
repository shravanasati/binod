package main

import (
	"fmt"
	"sort"
	"sync"
)

type Player struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Binods   int    `json:"binods"`
}

type playerDB struct {
	players []Player
	sync.Mutex
}

var playerdb playerDB

// addPlayer adds a new player to the database if the username doesnt exist and returns a
// boolean value of true if the username is unique.
func addPlayer(p *Player) (bool, string) {
	playerdb.Lock()
	defer playerdb.Unlock()

	for _, player := range playerdb.players {
		if player.Username == p.Username {
			return false, "This username already exists! Try another one."
		}
	}

	playerdb.players = append(playerdb.players, *p)

	return true, "Welcome " + p.Username + "!"
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

func updatePlayer(p *Player) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.Username == p.Username {
			playerdb.players[i].Binods = p.Binods
			return true
		}
	}

	return false
}

func deletePlayer(p *Player) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.Username == p.Username && player.Password == p.Password {
			playerdb.players = append(playerdb.players[:i], playerdb.players[i+1:]...)
			return true
		}
	}

	return false
}
