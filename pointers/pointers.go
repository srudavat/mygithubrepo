package pointers

import (
	"fmt"
	//"example.com/go-learn/dice"
)

type Players struct {
	health   int
	attack   int
	strength int
	name     string
}

func Combat() {

	fmt.Println("This is the code for Pointers")

	shiva := Players{name: "shiva", health: 70, attack: 10, strength: 40}
	marvin := Players{name: "marvin", health: 90, attack: 20, strength: 30}
	Game(&shiva, &marvin)

}

func Game(player1 *Players, player2 *Players) {

	var playerToattack *Players
	var playerTodefend *Players

	//Findout the player with lesser strength and make him the player who attacks.

	if player1.strength < player2.strength {
		playerToattack = player1
		playerTodefend = player2
	} else {
		playerToattack = player2
		playerTodefend = player1
	}
	fmt.Printf("Player to attack: %v\n", playerToattack.name)
	fmt.Printf("Player to defend: %v\n", playerTodefend.name)
	//v := dice.Rolldice()
	//player1.strength = player1.strength + v
	//player2.attack = player2.attack - v
}
