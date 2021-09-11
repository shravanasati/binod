package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// Player struct represents a player in the database.
type Player struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Binods   int    `json:"binods"`
}

// playerDB struct represents the database of players, and has a mutex associated with it.
type playerDB struct {
	players []Player
	sync.Mutex
}

// global player database
var playerdb playerDB

// global database cache
var dbCache = cache.New(5*time.Minute, 5*time.Minute)

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

// getLeaderboardData returns a slice of players sorted by binods.
func getLeaderBoardData() []Player {
	// t1 := time.Now()
	items, found := dbCache.Get("leaderboard")
	if found {
		// fmt.Println("Leaderboard cache found.")
		// fmt.Println("yes cache", time.Since(t1))
		return items.([]Player)
	}

	// t2 := time.Now()
	playerdb.Lock()
	defer playerdb.Unlock()

	sort.Slice(playerdb.players, func(i, j int) bool {
		return playerdb.players[i].Binods > playerdb.players[j].Binods
	})

	// fmt.Println("Setting cache...")
	dbCache.Set("leaderboard", playerdb.players, cache.DefaultExpiration)
	// fmt.Println("no cache", time.Since(t2))
	return playerdb.players
}

// updatePlayer updates a player data in the database.
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

// deletePlayer deletes a player from the database.
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

// getPlayer returns a player from the database.
func getPlayer(username string) (*Player, error) {
	data, found := dbCache.Get("player-" + username)
	if found {
		return data.(*Player), nil
	}

	playerdb.Lock()
	defer playerdb.Unlock()

	for _, player := range playerdb.players {
		if player.Username == username {
			dbCache.Set("player-"+username, &player, time.Minute)
			return &player, nil
		}
	}

	return &Player{}, fmt.Errorf("Player with username '%s' not found.", username)
}
