package main

import (
	"os"

	"github.com/Matts-vdp/terminal/ter"
)

func main() {
	l := make([]string, 2)
	term := ter.InitTerminal(os.Stdout, l, true)
	defer term.Close()
	term.UpdateLine(0, "Start")
	term.UpdateLine(1, "time passed")
}
