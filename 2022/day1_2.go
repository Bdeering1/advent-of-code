package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("day1/input.txt")
	if err != nil { panic(err) }
	defer f.Close()
	
	calories, max := 0, make([]int, 3)
	s := bufio.NewScanner(f)

	for s.Scan() {
		val, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			for i := 0; i < 3; i++ {
				if calories > max[i] {
					max[i] = calories
					sort.Ints(max)
					break
				}			
			}
			calories = 0
		}
		calories += int(val)
	}

	fmt.Println(max)
	fmt.Println(max[0] + max[1] + max[2])
}

