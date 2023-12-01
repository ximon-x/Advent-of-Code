package stuff

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var points map[string]int

func calcPoints(op string, cp string) int {
	s := points[cp]

	//  Rock        =>  A
	//  Paper       =>  B
	//  Scissors    =>  C

	//  Lose        =>  X
	//  Draw        =>  Y
	//  Win         =>  Z

	if op == "A" {
		if cp == "X" {
			s += points["C"]
		}

		if cp == "Y" {
			s += points["A"]
		}

		if cp == "Z" {
			s += points["B"]
		}
	}

	if op == "B" {
		if cp == "X" {
			s += points["A"]
		}

		if cp == "Y" {
			s += points["B"]
		}

		if cp == "Z" {
			s += points["C"]
		}

	}

	if op == "C" {
		if cp == "X" {
			s += points["B"]
		}

		if cp == "Y" {
			s += points["C"]
		}

		if cp == "Z" {
			s += points["A"]
		}

	}

	return s
}

func RockPaperScissors() (int, error) {
	absPath, err := filepath.Abs("../goony/assets/rock_paper_scissors_input.txt")
	if err != nil {
		fmt.Printf("Could not get asset.")
		return 0, err
	}

	file, err := os.ReadFile(absPath)
	data := string(file)

	chunks := strings.Split(data, "\n")

	points = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,

		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	score := 0

	for _, v := range chunks {
		if len(v) >= 2 {
			chars := strings.Split(v, "")
			score += calcPoints(chars[0], chars[2])
		}
	}

	return score, nil
}
