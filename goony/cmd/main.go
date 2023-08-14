package main

import (
	"fmt"
	"goony/stuff"
)

func main() {
	answer, err := stuff.RockPaperScissors()
	if err != nil {
		return
	}

	fmt.Println(answer)
}
