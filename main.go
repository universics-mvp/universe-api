package main

import (
	"log"

	"main/cmd"
)

func main() {
	err := cmd.StartApp()
	if err != nil {
		log.Fatal(err)
	}
}
