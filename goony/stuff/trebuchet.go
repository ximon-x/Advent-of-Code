package stuff

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func Trebutchet() (int, error) {
	absPath, err := filepath.Abs("../goony/assets/trebutchet_input.txt")
	if err != nil {
		fmt.Printf("Could not get asset.")
		return 0, err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Printf("Could not get file at %s\n", absPath)
		return 0, err
	}

	file := string(data)
	chunks := strings.Split(file, "\n")

	sum := 0

	for _, chunk := range chunks {
		firstDigit := 0
		lastDigit := 0

		for _, v := range chunk {
			if unicode.IsNumber(v) {
				// Converting from Rune to int is weird in Go.
				firstDigit = int(v-'0') * 10
				break
			}
		}

		for i := len(chunk) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(chunk[i])) {
				lastDigit = int(rune(chunk[i] - '0'))
				break
			}
		}

		sum += firstDigit + lastDigit

	}

	return sum, nil
}
