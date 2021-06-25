package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"malachi/pkg/bowling"
)

var (
	flagHelp *bool = flag.Bool("h", false, "Print help text.")
)

func main() {

	flag.Parse()
	if *flagHelp {
		help()
		return
	}

	// For the sake of time, I'm only going to support one user.

	score := bowling.NewScore()

	// The loop doesn't begin until the user provides input, so we need to provide some prompts up front.
	prompt(score)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		input := scanner.Text()
		frame := bowling.ParseFrame(input)
		if frame == nil {
			fmt.Println("Input not recognized. Please try again.")
		} else {
			score = score.AddFrame(*frame)
		}

		present(score)

		if len(score) == 10 {
			break
		}

		prompt(score)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func present(score bowling.Score) {
	fmt.Printf("Completed Frame: %d\n", len(score))
	pins := score.Pins()
	if pins == nil {
		fmt.Printf("Current Score: empty\n")
	} else {
		fmt.Printf("Current Score: %d\n", *pins)
	}
}

func help() {
	fmt.Print(`
This application allows the user to enter frames, and the application will 
respond with the current score (if calculable) until the end of the game.

When entering frame data, you may enter 'X' or '10' for a strike. Otherwise, it
is required that you enter comma-separated values to represent the first and
second throws, such as '2,3'. If you throw a spare, you may enter the numeric
values, such as '3,7', or the spare shorthand, like '3,/'.

The final (tenth) frame may include a third value. If the final frame represents
a spare, the third value will represent the fill. If the final frame represents 
a strike, the second a third values will represent the final throws to calculate
the value of the strike.
`)
}

func prompt(score bowling.Score) {
	fmt.Printf("Please enter frame #%d.\n", len(score)+1)
}
