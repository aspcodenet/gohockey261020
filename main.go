package main

// I go finns ALLT för att bygga en modern applikation
//  - multithreaded - goroutines
//  - webserver - ta emot http   <---- tredjepartsmodul Go Gin
// -  json, web

import (
	"fmt"

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

func main() {


	data.Init() // Initialize   players slice with sample data

	r := gin.Default()
	r.GET("/hej.html", hejHtmlHandler)
	r.GET("/api/players", getAllPlayersHandler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
}
