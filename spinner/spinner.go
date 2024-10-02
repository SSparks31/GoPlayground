package spinner

import (
	"fmt"
	"time"
)

type Spinner struct {
	frames    []string
	frameRate int
	messageCh chan string
	msg       string
	running   bool
}

func New() *Spinner {
	s := &Spinner{
		frames:    []string{"|", "/", "-", "\\"},
		frameRate: 7,
		messageCh: make(chan string),
		msg:       "",
		running:   false,
	}
	return s
}

func (s *Spinner) GetFrameRate() int {
	return s.frameRate
}

func (s *Spinner) Start() {
	s.running = true
	ticker := time.NewTicker(time.Millisecond * time.Duration(float64(1000)/float64(s.frameRate)))
	for {
		for _, frame := range s.frames {
			fmt.Printf("%s %s ", frame, s.msg)
		loop:
			for {
				if !s.running {
					fmt.Println("")
					return
				}
				select {
				case <-ticker.C:
					break loop
				case s.msg = <-s.messageCh:
					fmt.Print("\033[2K\r")
					fmt.Printf("%s %s ", frame, s.msg)
				}
			}
			fmt.Print("\033[2K\r")
		}
	}
}

func (s *Spinner) UpdateMessage(msg string) {
	s.messageCh <- msg
}

func (s *Spinner) Stop() {
	s.running = false
}
