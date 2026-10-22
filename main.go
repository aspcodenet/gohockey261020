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
	c.String(200, "<h1>Hej</h1><ul><li>Foppa</li><li>Zäta</li></ul>")
}

func getAllPlayersHandler(c *gin.Context) {
	players := data.GetAllPlayers()
	c.JSON(200, players)
}

func main() {

	//p:= data.New("Stefan", 12)

	data.Init() // Initialize   players slice with sample data
	r := gin.Default()
	r.GET("/hej.html", hejHtmlHandler)
	r.GET("/api/players", getAllPlayersHandler)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
	// //a := data.GetAllPlayers()
	// s := "Stefan"
	// // if s.Contains("Ste")
	// if strings.Contains(s, "Ste") {
	// 	fmt.Println("Found 'Ste' in the string!")
	// }

	// // Employee employee = new Employee();
	// // int money = emp.CalculateSalary();
	// // player = new Player();
	// // if (player.IsCool())

	// for _, player := range data.GetAllPlayers() {
	// 	//if data.IsCoolVersion1(player) {
	// 	if player.IsCool() {
	// 		fmt.Println(player.Name, player.Team, player.Number, player.Age, "is cool!")
	// 	} else {
	// 		fmt.Println(player.Name, player.Team, player.Number, player.Age, "is not cool.")
	// 	}
	//	}
}
