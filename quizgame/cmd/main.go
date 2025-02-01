package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/FedericoBarberon/Go-Exercises/quizgame"
)

const defaultPath = "problems.csv"
const defaultTimeLimitSec = 30

func main() {
	pathFlag := flag.String("csv", defaultPath, `a csv file in the format of 'question,answer'`)
	timeLimitSecFlag := flag.Int("limit", defaultTimeLimitSec, "the time limit for the quiz in seconds")
	flag.Parse()

	if *timeLimitSecFlag < 0 {
		fmt.Printf("invalid time limit")
		os.Exit(1)
	}

	problems, err := quizgame.GetProblemsFromFS(os.DirFS("./"), *pathFlag)

	if err != nil {
		fmt.Printf("error getting the data from file: %v", err)
		os.Exit(1)
	}

	questionAsker := quizgame.NewStdQuestionAsker()
	timer := quizgame.RealTimer{}
	timeLimit := time.Duration(*timeLimitSecFlag) * time.Second

	game := quizgame.NewGame(problems, questionAsker, timer, timeLimit)
	game.Play()
}
