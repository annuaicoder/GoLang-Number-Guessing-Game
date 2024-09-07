package main

import (
	"fmt"
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Guess a number between 1 and 100: ")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have",10-guesses,"guesses left")
		fmt.Printf("Make a guess: ")
		input,err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess,err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Too low!!!")
		} else if guess > target {
			fmt.Println("Too High!!!!!")
		} else {
			success = true
			fmt.Println("Good job! You gussed the number!")
			break
		}
	}
	if !success {
		fmt.Println("You could not guess the number! It was",target)
	}
}