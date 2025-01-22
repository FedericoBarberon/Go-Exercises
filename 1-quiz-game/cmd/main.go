package main

import (
	"flag"
	"fmt"
	"os"
	"quizgame"
)

const defaultPath = "problems.csv"

func main() {
	pathFlag := flag.String("csv", defaultPath, `a csv file in the format of 'question,answer'`)
	flag.Parse()

	qa, err := quizgame.GetQAFromFS(os.DirFS("./"), *pathFlag)

	if err != nil {
		fmt.Printf("error getting the data from file: %v", err)
		os.Exit(1)
	}

	game := quizgame.NewGame(qa, quizgame.NewStdQuestionAsker())
	game.Play()
}
