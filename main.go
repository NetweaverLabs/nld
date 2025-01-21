package main

import (
	"encoding/gob"
	"log"

	"github.com/NetweaverLabs/nld/daemon"
	"github.com/NetweaverLabs/types"
)

func init() {
	gob.Register(types.User{})
}

func main() {
	d, err := daemon.NewDaemon()
	if err != nil {
		log.Fatal(err)
	}
	d.Start()
}
