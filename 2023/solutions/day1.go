package solutions

import (
	"strings"
)

func (Solutions) Day1_1(input []string) int {
	sum := 0

	for _, line := range input {
		var digits []int
		for _, c := range line {
			dig := int(c - '0')
			if dig <= 9 {
				digits = append(digits, dig)
			}
		}
		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum
}

func (Solutions) Day1_2(input []string) int {
	patterns := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	sum := 0

	for _, line := range input {
		var digits []int
		var letters string

		for _, c := range line {
			dig := int(c - '0')
			if dig <= 9 {
				digits = append(digits, dig)
				letters = ""
			} else {
				letters += string(c)
				for idx, pat := range patterns {
					if strings.Contains(letters, pat) {
						digits = append(digits, idx+1)
						letters = letters[len(letters)-1:]
					}
				}
			}
		}
		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum
}
