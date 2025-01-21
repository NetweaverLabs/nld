package main

import (
	"log"

	"github.com/NetwaeversLab/nld/daemon"
)

func main() {
	d, err := daemon.NewDaemon()
	if err != nil {
		log.Fatal(err)
	}
	d.Start()
}
