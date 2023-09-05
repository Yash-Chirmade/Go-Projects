package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type player struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Team  string `json:"Team"`
	Score int    `json:"Score"`
}

var players []player
var nextID = 1

func createPlayer(c *gin.Context) {
	var p player
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.ID = nextID
	nextID++
	players = append(players, p)
	c.JSON(http.StatusCreated, p)
}

func getPlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	for _, p := range players {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
}

func updatePlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	var updatedPlayer player
	if err := c.ShouldBindJSON(&updatedPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range players {
		if p.ID == id {
			updatedPlayer.ID = id
			players[i] = updatedPlayer
			c.JSON(http.StatusOK, updatedPlayer)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
}

func deletePlayer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	for i, p := range players {
		if p.ID == id {
			players = append(players[:i], players[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
}

func getAllPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, players)
}

func main() {
	r := gin.Default()

	players = append(players, player{ID: nextID, Name: "Seth", Team: "MI", Score: 88})
	nextID++
	players = append(players, player{ID: nextID, Name: "Yash", Team: "CSK", Score: 100})
	nextID++
	players = append(players, player{ID: nextID, Name: "Karl", Team: "RCB", Score: 34})
	nextID++

	r.POST("/players", createPlayer)
	r.GET("/players", getAllPlayers)
	r.GET("/players/:id", getPlayer)
	r.PUT("/players/:id", updatePlayer)
	r.DELETE("/players/:id", deletePlayer)

	fmt.Println("Starting server at port 8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
