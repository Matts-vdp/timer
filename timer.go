package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Matts-vdp/timer/save"
)

// waits until enter in input
func waitRead() chan string {
	ch := make(chan string)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ch <- str
	}()
	return ch
}

// updates displayed time
func updateTime(now time.Time) {
	pas := time.Since(now)
	s := save.Print(pas)
	fmt.Print("\r", s)
}

// used to display al stored times
func getTimes() string {
	sav := save.Load()
	str := sav.String()
	return str
}

// updates screen until enter is pressed
func mainloop(now time.Time) {
	ch := waitRead()
	for {
		select {
		case <-ch:
			fmt.Println()
			return
		case <-time.After(time.Second / 5):
			updateTime(now)
		}
	}
}

func askMessage(now time.Time) {
	fmt.Print(">>")
	info := <-waitRead()
	save.Save(now, info)
}

func main() {
	lg := flag.Bool("l", false, "give logging info")
	flag.Parse()
	if *lg {
		fmt.Println(getTimes())
		return
	}
	now := time.Now()
	fmt.Println(now.Format("15:04:05"))
	mainloop(now)
	askMessage(now)
}
