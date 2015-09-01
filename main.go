package main

import (
	"github.com/elastic/topbeat/beat"
)

// You can overwrite these, e.g.: go build -ldflags "-X main.Version 1.0.0-beta3"
var Version = "1.0.0-beta2"
var Name = "topbeat"

func main() {

	tb := &Topbeat{}

	b := beat.NewBeat(Name, Version, tb)

	b.CommandLineSetup()

	b.LoadConfig()

	b.Run()

}
