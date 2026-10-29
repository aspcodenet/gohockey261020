package data

import (
	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

// koirrekt att det heter struct och inte class
// struct is a composite data type that groups together variables under a single name
type HockeyPlayer struct {
	Id     int
	Name   string `json:"name"` // struct tags
	Team   string
	Number int `json:"Jersey"` // annotera att den ska heta Jersey i json
	Age    int
}

func New(name string, age int) HockeyPlayer {
	p := HockeyPlayer{
		Name: name,
		Age:  age,
	}
	return p
}

func IsCoolVersion1(player HockeyPlayer) bool {
	if player.Age >= 54 {
		return true
	}
	return false
}

func (player HockeyPlayer) IsCool() bool {
	if player.Age >= 54 {
		return true
	}
	return false
}

var db *gorm.DB

// slice med alla hockeyspelare

func Init() {
	db, _ = gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	db.AutoMigrate(&HockeyPlayer{}) // CREATE TABLE om det inte finns

	// Seedning
	var antal int64
	db.Model(&HockeyPlayer{}).Count(&antal) // select count(*) from HockeyPlayer
	if antal == 0 {

		db.Create(&HockeyPlayer{Name: "Foppa", Team: "Modo", Number: 21, Age: 53}) // insert into

		db.Create(&HockeyPlayer{Name: "Foppa2", Team: "Modo", Number: 22, Age: 23}) // insert into
		db.Create(&HockeyPlayer{Name: "Foppa3", Team: "Modo", Number: 23, Age: 33}) // insert into

	}

}

func DeletePlayer(player HockeyPlayer) (HockeyPlayer, error) { //
	err := db.Delete(&player).Error
	return player, err
}

func UpdatePlayer(player HockeyPlayer) (HockeyPlayer, error) { //
	err := db.Save(&player).Error
	return player, err
}

func CreatePlayer(player HockeyPlayer) (HockeyPlayer, error) { //
	err := db.Create(&player).Error
	return player, err
}

func GetSinglePlayer(id string) (HockeyPlayer, error) {
	var player HockeyPlayer

	err := db.First(&player, id).Error
	return player, err
}

func GetAllPlayers() []HockeyPlayer {
	var players []HockeyPlayer
	db.Find(&players)
	return players
}

// OOP nerd - construtor som säkerställer Valid state = mandatory fields
// getters och setters för att säkerställa Valid state = mandatory fields

// Golang sämre än andra språk vad gäller OOP !

// brukar det inte vara ??? constructor, private, public, getter, setter, toString, equals, hashCode
