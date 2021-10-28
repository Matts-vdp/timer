package main

import (
	"fmt"
	"os"
	"time"
	"timer/save"

	"github.com/Matts-vdp/terminal/ter"
)

func waitRead() chan bool {
	ch := make(chan bool)
	go func() {
		var str string
		fmt.Scanln(&str)
		ch <- true
	}()
	return ch
}

func updateTime(term *ter.TerminalOut, now time.Time) {
	pas := time.Since(now)
	term.UpdateLine(0, now.Format("15:04:05"))
	s := save.Print(pas)
	term.UpdateLine(1, s)
}

func mainloop(now time.Time) {

	term := ter.InitTerminal(os.Stdout, make([]string, 2), true)
	defer term.Close()
	ch := waitRead()
	for {
		select {
		case <-ch:
			return
		case <-time.After(time.Second / 5):
			updateTime(term, now)
		}
	}
}

func main() {
	now := time.Now()
	mainloop(now)
	save.Save(now)
}
