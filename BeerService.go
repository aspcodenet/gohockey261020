package main

func CanIBuyBeer(age int, location string, promille float32) bool {
	//FÅNIGT - men symboliserar "affärslogik"
	// affärslogik består pfta av massa olika cases
	if promille > 1.5 {
		return false
	}
	if age >= 18 && location == "krogen" {
		return true
	}
	if age >= 20 && location == "systemet" {
		return true
	}
	return false
}
