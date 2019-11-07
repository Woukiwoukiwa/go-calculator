package main

import (
	"log"
	"os"
)

func main() {
	log.Print("Hello")
	argsWithoutProg := os.Args[1]
	log.Print(argsWithoutProg)
}
