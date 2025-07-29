package main

import (
	"log"
	"os"

	xz "github.com/takanoriyanagitani/go-xzcat"
)

func main() {
	if err := xz.Xz(os.Stdin, os.Stdout); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	return
}
