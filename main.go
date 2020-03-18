package main 

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	fmt.Println("Guessing Game")

	//generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) 

	var guess int 
	for {
	fmt.Println("Pick a number:")
	fmt.Scan(&guess)
	if guess != secretNumber {
		fmt.Println("Wrong number. Hint: It is an even number!")
	}else if guess > 10 {
		fmt.Println("We are far apart!")
	} else {
		fmt.Println("You got it Right!! You must be a genius")
		break
	}
	
	}
}