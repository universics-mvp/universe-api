package main

import "main/cmd"

func main() {

	err := cmd.StartApp()
	if err != nil {
		return
	}
}
