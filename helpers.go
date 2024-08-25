package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// TYPES
type Colors [3]uint8

type QuizGameStats struct {
	problems         []problem
	correctQuestions int
	totalQuestions   int
}

type problem struct {
	question string
	answer   string
}

// GENERAL HELPERS
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ceilDiv(a, b int) int {
	result := int(math.Ceil(float64(a) / float64(b)))
	return result
}

func generateRandN(minVal, maxVal int) int {
	return rand.Intn(maxVal-minVal) + minVal
}

func OpenProblemsCsv(file_path string) ([]problem, error) {
	file, err := os.Open(file_path)
	problems := []problem{}
	if err != nil {
		return problems, err
	}

	csvReader := csv.NewReader(file)
	problemsCSV, err := csvReader.ReadAll()
	if err != nil {
		return problems, err
	}

	for _, line := range problemsCSV {
		problems = append(
			problems, problem{line[0], strings.TrimSpace(line[1])})
	}

	return problems, nil
}

// FORMATTING FUNCTIONS

func ApplyColorToText(text string, color Colors) string {
	var Reset string = "\033[0m"
	var coloredText string = fmt.Sprintf("\033[38;2;%v;%v;%vm%v%v", color[0], color[1], color[2], text, Reset)
	return coloredText
}

func ResetScreen() {
	fmt.Println("\u001b[H")  // moves cursor to starting position
	fmt.Println("\u001b[0J") // clears cursor
}

// Routines
func StartTimer(timelimit int, quizGame *QuizGameStats) {
	timer := time.NewTimer(time.Second * time.Duration(flagTimeLimit))
	<-timer.C
	fmt.Println("Time limit exceeded")
	GameOverMessage(quizGame)
	os.Exit(1)
}

func ListenForExit(quizGame *QuizGameStats) {
	sigChan := make(chan os.Signal, 1)                      // making a channel of type os.Signal
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // we only send SIGINT and SIGTERM on the channel.
	<-sigChan

	GameOverMessage(quizGame)
	os.Exit(1)
}
