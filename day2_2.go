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
		line[0] -= 'A';
		
		if line[2] == 'X' { // lose
			score += int((line[0] + 5) % 3) + 1
			fmt.Printf("%c %c %d + 0\n", line[0] + 'A', line[2], int((line[0] + 5) % 3) + 1)
		} else if line[2] == 'Z' { // win
			score += int((line[0] + 4) % 3) + 6 + 1
			fmt.Printf("%c %c %d + 6\n", line[0] + 'A', line[2], int((line[0] + 4) % 3) + 1)
		} else { // draw
			score += int(line[0]) + 3 + 1
			fmt.Printf("%c %c %d + 3\n", line[0] + 'A', line[2], int(line[0]) + 1)
		}
	}
	fmt.Println(score)
}

