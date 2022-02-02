package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type ProblemType struct {
	problem string
	answer  string
}

type Quiz struct {
	problems []ProblemType
	score    int
}

func (qz *Quiz) playGame() {
	for _, el := range qz.problems {
		var userInput string
		fmt.Println(el.problem, "=>")
		fmt.Scanf("%s\n", &userInput)
		if userInput == el.answer {
			qz.score++
		}
	}
}

var (
	filePath *string
	timeout  *int
)

func init() {
	filePath = flag.String("csv", "quiz-cli/data.csv", "a csv file in the format of questions, answer")
	timeout = flag.Int("time", 30, "the time limit for the quiz in seconds")
	flag.Parse() // this would parse all the flags
}

func readCSV() []ProblemType {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalln("Failed to open the CSV file:", *filePath) // os.Exit(1)
	}
	defer file.Close()

	records, _ := csv.NewReader(file).ReadAll()
	problems := make([]ProblemType, len(records))
	for i, record := range records {
		problems[i] = ProblemType{problem: record[0], answer: record[1]}
	}
	return problems
}

func main() {
	questions := readCSV()
	qz := Quiz{questions, 0}

	timeOutCh := time.After((time.Duration(*timeout) * time.Second))
	resultCh := make(chan Quiz)

	go func() {
		qz.playGame()
		resultCh <- qz
	}()

	select {
	case <-resultCh:
	case <-timeOutCh:
		fmt.Println("TIMEOUT!!!")
	}

	fmt.Printf("Your Score %d out of %d.\n", qz.score, len(questions))
}
