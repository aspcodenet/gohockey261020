package main

// I go finns ALLT för att bygga en modern applikation
//  - multithreaded - goroutines
//  - webserver - ta emot http   <---- tredjepartsmodul Go Gin
// -  json, web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {

	data.Init() // Initialie   players slice with sample data

	r := gin.Default()
	r.GET("/hej.html", hejHtmlHandler)
	r.GET("/api/players", getAllPlayersHandler)    // lista alla
	r.GET("/api/players/:id", getSinglePlayerById) // hämta EN specifik
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
}
