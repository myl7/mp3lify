package main

import (
	"flag"
	"github.com/myl7/mp3lify"
	"log"
)

func main() {
	addr := flag.String("addr", "", "")
	flag.Parse()

	if err := mp3lify.Listen(*addr); err != nil {
		log.Fatalln(err)
	}

	log.Println("Listening on ", *addr)
}
