package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// FLAGS
var flagFilePath string
var flagShowHelp bool
var flagShuffle bool
var flagTimeLimit int

func showHelpMessage() {
	fmt.Println("Usage: QuizGame [options]")
	fmt.Println()
	fmt.Println("Options:")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf(" -%s\t%s\n", f.Name, f.Usage)
	})
}

func CreateQuiz() QuizGameStats {
	problems, err := OpenProblemsCsv(flagFilePath)
	if err != nil {
		log.Fatalf("Error parsing csv: %v\n", err)
	}
	return QuizGameStats{
		problems:         problems,
		correctQuestions: 0,
		totalQuestions:   len(problems),
	}

}

func init() {
	flag.StringVar(&flagFilePath, "file", "./problems.csv", "Specify the location of the file")
	flag.StringVar(&flagFilePath, "f", "./problems.csv", "Specify the location of the file")
	flag.BoolVar(&flagShowHelp, "help", false, "Show's help menu")
	flag.BoolVar(&flagShuffle, "s", false, "Shuffle questions")
	flag.IntVar(&flagTimeLimit, "t", 30, "Set time limit in seconds")
	flag.IntVar(&flagTimeLimit, "time", 30, "Set time limit in seconds")

	flag.Parse()
	if flagShowHelp {
		showHelpMessage()
		os.Exit(1)
	}
}

func main() {

	quizGame := CreateQuiz()

	println("\u001b[=19H")
	ResetScreen()

	// fmt.Println("Starting game...")
	StartGameMessage()
	timer := time.NewTimer(time.Second * time.Duration(flagTimeLimit))
	// go StartTimer(flagTimeLimit, &quizGame) // timer
	go ListenForExit(&quizGame) // checks for forced exit

	if flagShuffle {
		PlayShuffled(&quizGame, timer)
	} else {
		PlayInOrder(&quizGame, timer)
	}

	GameOverMessage(&quizGame)
}
