package main

import (
	"bearded/internal/app/commander"
	"flag"
	"fmt"
	"log"
)

func main() {
	var restore bool

	flag.BoolVar(&restore, "r", false, "defines if to restore info from dump or not")
	flag.Parse()

	fmt.Println(restore)
	cmd, err := commander.New(restore)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Run()
}
