package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
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

func askQuestion(questionIdx int, question Question, answerCh chan string) {
	fmt.Printf("Problem #%d: %s = ", questionIdx+1, question.question)
	var answer string
	fmt.Scanln(&answer)
	answerCh <- answer
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
	questions := readQuestions(r)
	correct := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

quizGame:
	for i, question := range questions {
		answerCh := make(chan string)
		go askQuestion(i, question, answerCh)
		select {
		case <-timer.C:
			fmt.Println("Time out!")
			break quizGame
		case answer := <-answerCh:
			if answer == question.answer {
				correct += 1
			}
		}
	}

	printResults(correct, len(questions))
}
