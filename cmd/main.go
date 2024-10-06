package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/just-umyt/just-game/internal/about"
	"github.com/just-umyt/just-game/internal/start"
)

func main() {
	newGame := flag.Bool("start", false, "to start a game")
	aboutGame := flag.Bool("about", false, "about the rule")

	flag.Parse()

	switch {
	case *newGame:
		if pro := flag.Arg(0); pro == "pro" {
			start.Start(true)
		} else {
			start.Start(false)
		}
	case *aboutGame:
		about.About()
	default:
		fmt.Fprintln(os.Stderr, errors.New("invalid command"))
		os.Exit(1)
	}

}
