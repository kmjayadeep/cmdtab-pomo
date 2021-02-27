package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rwxrob/conf-go"
)

var config = conf.New()

func usage() {
	fmt.Println("pomo [clear|start]")
}

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]

	if len(args) == 0 {
		show()
		return
	}

	switch args[0] {
	case "clear":
		config.SetSave("pomo.up", "")
	case "dur":
		// TODO validate the duration
		config.SetSave("pomo.dur", args[1])
	case "start":
		// TODO detect optional duration argument
		s := config.Get("pomo.dur")
		if len(args) > 1 {
			s = args[1]
		}
		if s == "" {
			s = "25m"
		}
		config.Set("pomo.dur", s)
		dur, err := time.ParseDuration(s)
		if err != nil {
			panic(err)
		}
		up := time.Now().Add(dur).Format(time.RFC3339)
		config.SetSave("pomo.up", up)
	default:
		usage()
		return
	}
}

func show() {
	u := config.Get("pomo.up")
	if u != "" {
		up, err := time.Parse(time.RFC3339, u)
		if err != nil {
			panic(err)
		}
		if up.After(time.Now()) {
			fmt.Printf(" %v\n", up.Sub(time.Now()).Round(time.Second))
		} else {
			fmt.Printf(" %v\n", time.Now().Sub(up).Round(time.Second))
		}
	}
}
