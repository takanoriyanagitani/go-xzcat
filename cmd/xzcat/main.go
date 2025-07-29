package main

import (
	"log"
	"os"

	xz "github.com/takanoriyanagitani/go-xzcat"
)

func main() {
	err := xz.Xzcat(os.Stdin, os.Stdout)
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
}
