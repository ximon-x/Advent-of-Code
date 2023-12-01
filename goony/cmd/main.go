package main

import (
	"fmt"
	"goony/stuff"
)

func main() {
	answer, err := stuff.Trebutchet()
	if err != nil {
		return
	}

	fmt.Println(answer)
}
