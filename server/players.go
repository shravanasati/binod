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

// addPlayer adds a new player to the database if the username doesn't exist and returns a
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

func getLeaderBoardData() []Player {
	playerdb.Lock()
	defer playerdb.Unlock()

	sort.Slice(playerdb.players, func(i, j int) bool {
		return playerdb.players[i].Binods > playerdb.players[j].Binods
	})

	return playerdb.players
}

func updatePlayer(p *Player) bool {
	playerdb.Lock()
	defer playerdb.Unlock()

	for i, player := range playerdb.players {
		if player.Username == p.Username && player.Password == p.Password {
			playerdb.players[i] = *p
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

func getPlayer(username string) (*Player, error) {
	playerdb.Lock()
	defer playerdb.Unlock()

	for _, player := range playerdb.players {
		if player.Username == username {
			return &player, nil
		}
	}

	return &Player{}, fmt.Errorf("Player with username '%s' not found.", username)
}
