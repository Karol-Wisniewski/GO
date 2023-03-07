package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	r "math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func game() {

	var n int = r.Intn(100)

	attempts := 0

	for {
		attempts++

		fmt.Print("Welcome to guessing game! To exit - enter 'koniec'. What's your guess? ")

		//Reading number or endgame info from player (string)

		in := bufio.NewReader(os.Stdin)
		guessStr, errRead := in.ReadString('\n')

		guessStr = strings.TrimFunc(guessStr, func(r rune) bool {
			return unicode.IsSpace(r)
		})

		if guessStr == "koniec" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		//Converting player's input to int

		guess, errConversion := strconv.Atoi(guessStr)

		if errRead != nil {
			fmt.Println("Error during data input!")
			return
		}

		if errConversion != nil {
			fmt.Println("Error during conversion!")
			return
		}

		//Checking game conditions

		if guess < n {
			fmt.Println("Your number is too small. Try again.")
		} else if guess > n {
			fmt.Println("Your number is too big. Try again.")
		} else {
			fmt.Println("You guessed the number after", attempts, "attempts!")
			fmt.Println("Well done, enter your name for ranking: ")

			//Reading player's name

			inName := bufio.NewReader(os.Stdin)
			name, errName := inName.ReadString('\n')

			if errName != nil {
				fmt.Println("Error during name input!")
				return
			}

			name = strings.TrimFunc(name, func(r rune) bool {
				return unicode.IsSpace(r)
			})

			//Open scores.txt if exists, if not create it and write player's name and score into it

			file, err := os.OpenFile("scores.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

			if err != nil {
				fmt.Println(err)
				return
			}

			defer file.Close()

			attemptsStr := strconv.Itoa(attempts)

			attemptsStr = strings.TrimFunc(attemptsStr, func(r rune) bool {
				return unicode.IsSpace(r)
			})

			_, err = file.WriteString(name + " " + attemptsStr + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}

			//Check if player wants to play again

			fmt.Println("Wanna play again? [y/n]")

			in := bufio.NewReader(os.Stdin)
			answer, err := in.ReadString('\n')

			if err != nil {
				fmt.Println("Error during data input.")
				return
			}

			answer = strings.TrimFunc(answer, func(r rune) bool {
				return unicode.IsSpace(r)
			})

			if answer == "y" {
				fmt.Println("Restarting the game...")
				game()
			} else {

				content, err := ioutil.ReadFile("scores.txt")
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}

				// Split the contents of the file into separate lines
				lines := strings.Split(string(content), "\n")

				// Define a slice to store the scores
				scores := make([][2]string, 0)

				// Loop over each line and split it into a name and score
				for _, line := range lines {
					parts := strings.Split(line, " ")
					if len(parts) != 2 {
						// Skip lines that don't contain a name and score
						continue
					}

					name := parts[0]
					score := parts[1]

					// Append the name and score to the slice of scores
					scores = append(scores, [2]string{name, score})
				}

				// Sort the slice of scores by the lowest to highest score
				sort.Slice(scores, func(i, j int) bool {
					scoreI := 0
					scoreJ := 0
					fmt.Sscanf(scores[i][1], "%d", &scoreI)
					fmt.Sscanf(scores[j][1], "%d", &scoreJ)
					return scoreI < scoreJ
				})

				// Print the sorted scores
				for _, score := range scores {
					fmt.Println("Name:", score[0])
					fmt.Println("Score:", score[1])
				}

				fmt.Println("Exiting...")
				os.Exit(0)
			}
		}
	}
}

func main() {

	fi, err := os.Stat("scores.txt")

	if err == nil && fi.Size() > 0 {
		os.Truncate("scores.txt", 0)
	}

	game()
}
