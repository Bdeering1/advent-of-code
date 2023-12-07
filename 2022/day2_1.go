package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day2/input.txt")
	if err != nil { panic(err) }
	defer f.Close()
	
	score := 0
	s := bufio.NewScanner(f);	
	for s.Scan() {
		if err != nil { break }
		line := []rune(s.Text())
		line[2] -= 23

		score += int(line[2] - 'A') + 1
		if line[0] == line[2] {
			score += 3
		} else if (line[2] == 'A' && line[0] == 'C') || (line[0] < line[2] && !(line[0] == 'A' && line[2] == 'C')) {
			score += 6
		}
	}
	fmt.Println(score)
}

