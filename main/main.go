package main

import (
	"flag"
	"github.com/myl7/mp3lify"
	"log"
)

func main() {
	addr := flag.String("addr", ":8000", "")
	flag.Parse()

	log.Println("Listening on ", *addr)
	if err := mp3lify.Listen(*addr); err != nil {
		log.Fatalln(err)
	}
}
