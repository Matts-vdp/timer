package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"timer/save"

	"github.com/Matts-vdp/terminal/ter"
)

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

func updateTime(term *ter.TerminalOut, now time.Time) {
	pas := time.Since(now)
	term.UpdateLine(0, now.Format("15:04:05"))
	s := save.Print(pas)
	term.UpdateLine(1, s)
}

func GetTimes() string {
	sav := save.Load()
	str := sav.String()
	return str
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
	lg := flag.Bool("l", false, "give logging info")
	flag.Parse()
	if *lg {
		fmt.Println(GetTimes())
		return
	}
	now := time.Now()
	mainloop(now)
	fmt.Print(">>")
	info := <-waitRead()
	save.Save(now, info)
}
