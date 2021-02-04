package main

import "fmt"

type user struct {
	name    string
	age     int
	isAdmin bool
}

func gokuTest() {
	goku := user{
		name: "Goku",
		age:  16,
	}
	fmt.Println("Goku >>", goku)
	goku.isAdmin = true
	fmt.Println("Goku >> After ", goku)

	// ***

	berk := user{"Berk", 16, false}
	fmt.Println("Berk >>", berk)

	// Anonim struct'lar

	alex := struct {
		name         string
		age          int
		team         string
		isGoalkeeper bool
	}{"Alex De Souza", 36, "Fenerbahçe", false}

	fmt.Println("Alex =>", alex)
}

func main() {
	gokuTest()
}
