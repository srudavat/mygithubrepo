package main

import (
	"example.com/go-learn/pointers"
	"fmt"
)

type Player struct {
	name   string
	health int
}

func main() {
	//var world string
	//	var first float32
	//	var second float32
	//	var div1 float32
	//	var div2 float32
	//
	//	fmt.Println("Enter your first number:")
	//
	//	fmt.Scan(&first)
	//
	//	fmt.Println("Enter your second number:")
	//	fmt.Scan(&second)
	//
	//	fmt.Println("Enter Divide number one:")
	//
	//	fmt.Scan(&div1)
	//
	//	fmt.Println("Enter Divide number two:")
	//	fmt.Scan(&div2)
	//
	//	//inputString := "Welcome to 30 Days of Code!"
	//	//world = "Hello, World!"
	//	fmt.Println("This is the Sum of these numbers")
	//	var sum float32
	//	//sum = first + second
	//	//fmt.Println(world)
	//	sum = sumdecimal(first, second)
	//
	//	fmt.Printf("%.2f\n", sum)
	//
	//	fmt.Println("Thala for a reason")
	//
	//	fmt.Println("This is the division of two decimals")
	//
	//	var division float32
	//	division = divide(div1, div2)
	//
	//	fmt.Printf("%.2f\n", division)
	//
	//	//fmt.Println("Now lets us multiply numbers:")
	//	//
	//	//var multiply int
	//	//
	//	//multiply = multipynumbers(first, second)
	//	//
	//	//fmt.Printf("%d\n", multiply)
	//
	//	//hash := sha256.New()
	//	//fmt.Println(hash)
	//
	//}
	//
	//func sumnumbers(a int, b int) int {
	//	sum := a + b
	//	return sum
	//}
	//
	//func multipynumbers(c int, d int) int {
	//
	//	multiply := c * d
	//	return multiply
	//}
	//
	//func sumdecimal(e float32, f float32) float32 {
	//
	//	decimal := e + f
	//	return decimal
	//
	//}
	//
	//func divide(r float32, s float32) float32 {
	//
	//	divide := r / s
	//	return divide
	//}

	//var a int
	////var b *int
	//
	//a = 5
	////b = &a
	//
	//d := pointermultiplier(&a)

	//fmt.Printf("This is the value of Point %d\n", d)
	//
	//fmt.Printf("this is first value %d\n", a)
	//
	//fmt.Printf("this is second value %d\n", &b)
	//
	//fmt.Printf("this is third value %d\n", c)
	//
	//fmt.Printf("This is the 5 multiplied value: %d\n", c)
	//
	//fmt.Printf("This is the pointer multiplied value: %d\n", c)

	//var shiva Player
	//
	//shiva.name = "Ram"
	//shiva.health = 100
	//
	////var marvin Player
	//
	//marvin := Player{name: "sre", health: 50}
	//
	//fmt.Println(marvin)
	//r := dice.Rolldice()
	//fmt.Printf("%d\n", r)

	pointers.Combat()

}

func pointermultiplier(e *int) int {
	*e = *e * 5
	fmt.Printf("%v\n", *e)
	return *e
}

func srecapacity(player *Player) {
	damage := 2
	player.health -= damage
}

//Assignment: Create a GO program which will take two players who are intialiazed with health, attack and strength. Create a function called fight
//which will take two input parameters of two players. Create a Dice roll which will have number between 1 and 6. Multiply player1
//attack by the outcome of Dice roll. Multiply the same dice roll with the strength of player 2. Always a player with low strength roll attacks.
//	The game ends when the health if one of the players health goes down to 0.
