package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// PLAYING ORDER
func PlayInOrder(quizGame *QuizGameStats) {
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= quizGame.totalQuestions; i++ {
		QuestionHandler(quizGame, i, reader)
	}
}

func PlayShuffled(quizGame *QuizGameStats) {
	reader := bufio.NewReader(os.Stdin)
	questionsVisited := make(map[int]bool)

	for i := generateRandN(1, quizGame.totalQuestions+1); len(questionsVisited) < quizGame.totalQuestions; i = generateRandN(1, quizGame.totalQuestions+1) {
		if questionsVisited[i] {
			continue
		}
		questionsVisited[i] = true
		QuestionHandler(quizGame, i, reader)
	}
}

func QuestionHandler(quizGame *QuizGameStats, i int, reader *bufio.Reader) {
	var Reset = "\033[0m"
	var Yellow = "\033[33m"
	fmt.Printf(Yellow+"Question %v:"+Reset+" %v\n", i, quizGame.problems[i].question)

	// Readstring will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Err reading input: %v", err)
		return
	}
	input = strings.TrimSpace(strings.TrimSuffix(input, "\n")) // remove delimiter from string

	if strings.EqualFold(input, quizGame.problems[i].answer) {
		quizGame.correctQuestions++
	}
	ResetScreen()
}

func StartGameMessage() {
	// R, O, Y, G, B, I, V
	rainbow := []Colors{{255, 0, 0}, {255, 165, 0}, {255, 255, 0}, {0, 255, 0}, {0, 0, 255}, {75, 0, 130}, {238, 130, 238}}
	file, err := os.Open("./startgame.txt")
	if err != nil {
		log.Fatalf("Err :%v\n", err)
		return
	}
	scanner := bufio.NewScanner(file)
	startText := ""
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)
		div := ceilDiv(lineLen, 7)
		for i, step := 0, 0; i < 7; i, step = i+1, step+div {
			upperBound := min(step+div, lineLen)
			appendingText := fmt.Sprintf("%v", ApplyColorToText(line[step:upperBound], rainbow[i]))
			startText += appendingText
		}
		startText += "\n"
	}
	fmt.Println(startText)
	fmt.Print("Press enter to begin")
	var input string
	fmt.Scanf("%v\n", &input)

	time.Sleep(1 * time.Second)

	ResetScreen()
}

func GameOverMessage(quizGame *QuizGameStats) {
	ResetScreen()
	fmt.Println(ApplyColorToText("QUIZ END!", Colors{255, 165, 0}))
	fmt.Printf("Questions Correct: %v\n", ApplyColorToText(strconv.Itoa(quizGame.correctQuestions), Colors{0, 255, 0}))
	fmt.Printf("Total Questions: %v\n", ApplyColorToText(strconv.Itoa(quizGame.totalQuestions), Colors{0, 255, 255}))
}
