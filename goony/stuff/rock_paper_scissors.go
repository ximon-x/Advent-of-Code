package stuff

import (
	"os"
	"strings"
)

var points map[string]int

func calcPoints(op string, cp string) int {
	s := points[cp]

	//  Rock =>      A & X
	//  Paper =>     B & Y
	//  Scissors =>  C & Z

	if op == "A" {
		if cp == "X" {
			s += points["D"]
		} else if cp == "Y" {
			s += points["W"]
		} else if cp == "Z" {
			s += points["L"]
		}
	}

	if op == "B" {
		if cp == "X" {
			s += points["L"]
		} else if cp == "Y" {
			s += points["D"]
		} else if cp == "Z" {
			s += points["W"]
		}
	}

	if op == "C" {
		if cp == "X" {
			s += points["W"]
		} else if cp == "Y" {
			s += points["L"]
		} else if cp == "Z" {
			s += points["D"]
		}
	}

	return s
}

func RockPaperScissors() (int, error) {
	d, err := os.ReadFile("/home/simon/Projects/Personal/Advent-of-Code/goony/assets/rock_paper_scissors_input.txt")
	if err != nil {
		return 0, err
	}

	data := string(d)
	chunks := strings.Split(data, "\n")

	points = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,

		"L": 0,
		"D": 3,
		"W": 6,
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
