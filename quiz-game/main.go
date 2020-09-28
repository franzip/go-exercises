package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Question struct {
	question string
	answer   string
}

func (q Question) isValidAnswer(answer string) bool {
	return answer == q.answer
}

func readQuestions(r *csv.Reader) (questions []Question) {
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("An error occurred: ", err)
		}

		questions = append(questions, Question{question: record[0], answer: record[1]})
	}
	return
}

func printResults(correctAnswers int, totalQuestions int) {
	fmt.Println("------------------")
	fmt.Println("Results:")
	fmt.Printf("Number of correct answers: %d/%d", correctAnswers, totalQuestions)
}

func runQuiz(questions []Question) {
	var answer string
	correctAnswers := 0

	for idx, question := range questions {
		fmt.Printf("Question number %d: %s \n", idx+1, question.question)
		fmt.Printf("Your answer: ")
		fmt.Scanln(&answer)
		if question.isValidAnswer(answer) {
			correctAnswers += 1
		}
	}

	printResults(correctAnswers, len(questions))
}

var fileName string
var timeLimit int

func main() {
	flag.StringVar(&fileName, "f", "problems.csv", "-f <filename>.csv")
	flag.IntVar(&timeLimit, "t", 30, "-t 50 // 50 seconds")
	flag.Parse()

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	var questions []Question = readQuestions(r)

	fmt.Println("Welcome to the simplest quiz game ever. Press enter to start")

	fmt.Scanln()

	if err != nil {
		log.Fatal("An error occurred reading user input: ", err)
	}

	runQuiz(questions)
}
