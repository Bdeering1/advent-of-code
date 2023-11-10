package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	score := 0
	line_count := 0
	items := make(map[rune]int)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if err != nil {
			break
		}

		line := s.Text()
		group_num := line_count % 3

		fmt.Println(line)

		for _, c := range line {
			if items[c] == group_num {
				items[c]++
			}
			if items[c] == 3 {
				items = make(map[rune]int)
				fmt.Printf("%c ", c)
				if c >= 'a' { c -= 58 } // line up lower case letters below uppercase in ASCII
				score += int(c - 'A') + 27
				fmt.Printf("(%d)\n", int(c - 'A') + 27)
				break
			}
		}
		line_count++

		if group_num == 2 {
			fmt.Println()
		}
	}
	fmt.Println(score)
}

