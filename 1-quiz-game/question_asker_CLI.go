package quizgame

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

func (cli *QuestionAskerCLI) ShowScore(score, totalQuestions int) {
	cli.out.Write([]byte(fmt.Sprintf("You scored %d out of %d", score, totalQuestions)))
}

func (cli *QuestionAskerCLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
