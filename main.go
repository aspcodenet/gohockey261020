package main

// I go finns ALLT för att bygga en modern applikation
//  - multithreaded - goroutines
//  - webserver - ta emot http   <---- tredjepartsmodul Go Gin
// -  json, web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"systementor.se/gohockey261020/data"
)

func hejHtmlHandler(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", []byte("<h1>Hej</h1><ul><lia>Foppa</li><li>Zäta</li></ul>"))
}

func getAllPlayersHandler(c *gin.Context) {
	players := data.GetAllPlayers()
	c.JSON(200, players)
}

func getSinglePlayerById(c *gin.Context) {
	id := c.Param("id") // id kommer vara 2
	player, err := data.GetSinglePlayer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.IndentedJSON(http.StatusOK, player)
	}
	// hämta player med id 2 från db
	// get single player
}

func deletePlayerById(c *gin.Context) {
	id := c.Param("id")
	player, err := data.GetSinglePlayer(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		player.Id, _ = strconv.Atoi(id)
		data.DeletePlayer(player)
		c.IndentedJSON(http.StatusNoContent, "")
	}
}

func updatePlayerById(c *gin.Context) {
	id := c.Param("id")

	player, err := data.GetSinglePlayer(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		if err := c.BindJSON(&player); err != nil {
			return
		}
		player.Id, _ = strconv.Atoi(id)
		data.UpdatePlayer(player)
		c.IndentedJSON(http.StatusOK, player)
	}
}

func createNewPlayer(c *gin.Context) {
	// eftersom vi kör REST och JSON så är tanken - en client, typ
	//{
	//"name": "Stefan",
	//"Team": "Bla",
	//"Jersey": 11,
	//"Age": 11
	//}
	var err error
	var player data.HockeyPlayer

	if err := c.BindJSON(&player); err != nil { // det som kom in via HTTP POST (i formatet JSON) mappa om det till ett objekt
		return
	}
	player.Id = 0
	player, err = data.CreatePlayer(player)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, player)
	}

}

func main() {

	data.Init() // Initialie   players slice with sample data

	r := gin.Default()
	r.GET("/hej.html", hejHtmlHandler)
	r.GET("/api/players", getAllPlayersHandler)    // lista alla
	r.GET("/api/players/:id", getSinglePlayerById) // hämta EN specifik
	r.POST("/api/players", createNewPlayer)
	r.PUT("/api/players/:id", updatePlayerById)
	r.DELETE("/api/players/:id", deletePlayerById)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
}
