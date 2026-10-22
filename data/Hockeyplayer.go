package data

// koirrekt att det heter struct och inte class
// struct is a composite data type that groups together variables under a single name
type HockeyPlayer struct {
	Name   string
	Team   string
	Number int
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

// slice med alla hockeyspelare
var players []HockeyPlayer

func Init() {
	players = []HockeyPlayer{
		{Name: "John Doe", Team: "Team A", Number: 10, Age: 25},
		{Name: "Jane Smith", Team: "Team B", Number: 15, Age: 22},
		{Name: "Mike Johnson", Team: "Team A", Number: 8, Age: 28},
	}
}

func GetAllPlayers() []HockeyPlayer {
	return players // på tisdag = vi gör om denna så den hämtar ifrån databas
}

// OOP nerd - construtor som säkerställer Valid state = mandatory fields
// getters och setters för att säkerställa Valid state = mandatory fields

// Golang sämre än andra språk vad gäller OOP !

// brukar det inte vara ??? constructor, private, public, getter, setter, toString, equals, hashCode
