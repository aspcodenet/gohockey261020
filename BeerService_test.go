package main

//OBS: Svårt? Helt OK
// Vi är på en programmeringsteknisk NÖRD-nivå
//det finns inga krav på att ni ska lära er skriva tester
//jag kommer ge er tester ifall det är så
// det viktiga är inse att det är en del i PROCESSEN

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPromilleIsGreaterThan1_5ThenNotAllowed(t *testing.T) {
	// arrange
	promille := float32(1.6)
	age := 50
	location := "krogen"

	// act
	canBuy := CanIBuyBeer(age, location, promille)

	// assert
	assert.False(t, canBuy)
}

func TestWhenPromilleIsLessThan1_5AndAgeIs18AndLocationIsKrogenThenAllowed(t *testing.T) {
	// arrange
	promille := float32(1.4)
	age := 18
	location := "krogen"

	// act
	canBuy := CanIBuyBeer(age, location, promille)

	// assert
	assert.True(t, canBuy)
}

func TestWhenPromilleIsLessThan1_5AndAgeIs20AndLocationIsSystemetThenAllowed(t *testing.T) {
	// arrange
	promille := float32(1.4)
	age := 20
	location := "systemet"

	// act
	canBuy := CanIBuyBeer(age, location, promille)

	// assert
	assert.True(t, canBuy)
}

func TestWhenPromilleIsLessThan1_5AndAgeIs19AndLocationIsSystemetThenNotAllowed(t *testing.T) {
	// arrange
	promille := float32(1.4)
	age := 19
	location := "systemet"

	// act
	canBuy := CanIBuyBeer(age, location, promille)

	// assert
	assert.False(t, canBuy)
}

func TestWhenPromilleIsLessThan1_5AndAgeIs17AndLocationIsKrogenThenNotAllowed(t *testing.T) {
	// arrange
	promille := float32(1.4)
	age := 17
	location := "krogen"

	// act
	canBuy := CanIBuyBeer(age, location, promille)

	// assert
	assert.False(t, canBuy)
}
