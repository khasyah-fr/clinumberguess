package main

import (
	"fmt"
	"math/rand"
	"time"
)

var guess uint8 //guess possibility: 0 to 255

var tries uint8 = 3 //how many times player can guess

func triesHandler(tries uint8, i uint8) {
	attempts := uint8(tries - (i + 1))
	if attempts > 0 {
		fmt.Printf("You have %v tries \n \n", attempts)
	} else {
		fmt.Println("Sorry, no more tries")
	}
}

func main() {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	secretNum := uint8(random.Intn(10)) //change 10 to your desired range

	fmt.Println("Guess the number between 0 and 9") //change the string based on [0,rangeLimit)
	fmt.Printf("You have %v tries \n", tries)

	var i uint8

	for i = 0; i < tries; i++ {
		fmt.Print("Your guess: ")
		fmt.Scan(&guess)
		if guess > secretNum {
			fmt.Println("Too big")
			triesHandler(tries, i)
		} else if guess < secretNum {
			fmt.Println("Too small")
			triesHandler(tries, i)
		} else {
			fmt.Println("HOORAY!")
			fmt.Printf("You used %v tries", (i + 1))
			break
		}
	}
}
