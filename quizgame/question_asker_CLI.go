package quizgame

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const TimeOverPrefixMessage = "\nTime over: "

type QuestionAskerCLI struct {
	in  *bufio.Scanner
	out io.Writer
}

func NewQuestionAskerCLI(in io.Reader, out io.Writer) *QuestionAskerCLI {
	return &QuestionAskerCLI{
		in:  bufio.NewScanner(in),
		out: out,
	}
}

func NewStdQuestionAsker() *QuestionAskerCLI {
	return &QuestionAskerCLI{
		in:  bufio.NewScanner(os.Stdin),
		out: os.Stdout,
	}
}

func (cli *QuestionAskerCLI) AskQuestion(q string) (answer string) {
	q += ": "
	cli.out.Write([]byte(q))
	answer = cli.readLine()
	return
}

func (cli *QuestionAskerCLI) ShowScore(score, totalQuestions int, timeOver bool) {
	var message string
	if timeOver {
		message = TimeOverPrefixMessage
	}
	message += fmt.Sprintf("You scored %d out of %d", score, totalQuestions)

	cli.out.Write([]byte(message))
}

func (cli *QuestionAskerCLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
