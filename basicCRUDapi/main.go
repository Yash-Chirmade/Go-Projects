package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type player struct {
	Name  string `json:"Name"`
	Team  string `json:"Team"`
	Score int    `json:"Score"`
}

var players []player

func addPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var playerToAdd player
	_ = json.NewDecoder(r.Body).Decode(&playerToAdd)

	players = append(players, playerToAdd)
	json.NewEncoder(w).Encode(playerToAdd)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, player := range players {
		if player.Name == param["Name"] {
			fmt.Print(player)
			json.NewEncoder(w).Encode(player)
			break
		}
	}
}
func updateScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	var Player player
	for index, item := range players {
		if item.Name == param["Name"] {
			_ = json.NewDecoder(r.Body).Decode(&Player)
			ScoreInt, _ := strconv.Atoi(param["Score"])
			Player.Score = ScoreInt
			players[index] = Player
			return
		}
		http.NotFound(w, r)
	}
}
func deletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	for index, item := range players {
		if item.Name == param["Name"] {
			players = append(players[:index], players[index+1:]...)
			json.NewEncoder(w).Encode(players)
			return
		}
	}
}


func main() {
	r := mux.NewRouter()
	players = append(players, player{Name: "seth", Team: "MI", Score: 88})
	players = append(players, player{Name: "Yash", Team: "CSK", Score: 100})
	players = append(players, player{Name: "karl", Team: "RCB", Score: 34})
	r.HandleFunc("/addplayer", addPlayer).Methods("POST")
	r.HandleFunc("/updateScore/{Name}", updateScore).Methods("PUT")
	r.HandleFunc("/deletePlayer/{Name}", deletePlayer).Methods("DELETE")
	r.HandleFunc("/getplayer/{Name}", getPlayer).Methods("GET")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
