package quizgame

import "time"

type RealTimer struct{}

func (t RealTimer) StartTimer(timeLimit time.Duration) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		time.Sleep(timeLimit)
		close(ch)
	}()

	return ch
}
