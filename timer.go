package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Matts-vdp/terminal/ter"
)

func main() {
	l := make([]string, 2)
	term := ter.InitTerminal(os.Stdout, l, true)
	defer term.Close()
	term.UpdateLine(0, "Start met tellen")
	term.UpdateLine(1, "time passed")
	now := time.Now()
	for {
		pas := time.Since(now)
		term.UpdateLine(0, now.Format("15:04:05"))
		hr := int(pas.Hours())
		mn := int(pas.Minutes()) - int(pas.Hours())*60
		sc := int(pas.Seconds()) - int(pas.Minutes())*60
		s := fmt.Sprintf("%.2d:%.2d:%.2d", hr, mn, sc)
		term.UpdateLine(1, s)
		<-time.After(time.Second / 6)
	}
}
