package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day3/input.txt")
	if err != nil { panic(err) }
	defer f.Close()
	
	score := 0
	s := bufio.NewScanner(f);	
	for s.Scan() {
		if err != nil { break }

		items := make(map[rune]bool)
		line := s.Text()
		for _, c := range line[:len(line) / 2] {
			items[c] = true
		}
		for _, c := range line[len(line) / 2 : len(line)] {
			if items[c] { 
				if c >= 'a' { c -= 58 } // line up lower case letters below uppercase in ASCII
				score += int(c - 'A') + 27
				if int(c - 'A') + 27 < 1 || int(c - 'A') + 27 > 52 {
					fmt.Println("Output range error")
				}
				break
			}
		}
	}
	fmt.Println(score)
}

