package main

import (
	"time"

	"github.com/SSparks31/GoPlayground/spinner"
)

func main() {
	s := spinner.New()
	go s.Start()
	s.UpdateMessage("Spinner go brrrrrrrrrrrrrr")
	time.Sleep(time.Second * 5)
	s.UpdateMessage("I live in Spain without the A")
	time.Sleep(time.Second * 5)
	s.UpdateMessage("Stopping in five seconds")
	time.Sleep(time.Second * 1)
	s.UpdateMessage("Stopping in four seconds")
	time.Sleep(time.Second * 1)
	s.UpdateMessage("Stopping in three seconds")
	time.Sleep(time.Second * 1)
	s.UpdateMessage("Stopping in two seconds")
	time.Sleep(time.Second * 1)
	s.UpdateMessage("Stopping in one second")
	time.Sleep(time.Second * 1)
	s.UpdateMessage("Stopped")
	s.Stop()
	time.Sleep(time.Hour)
}
