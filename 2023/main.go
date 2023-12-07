package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "reflect"
    "github.com/bdeering1/advent-of-code/2023/solutions"
)

func main() {
    day := flag.Int("day", 0, "day to execute")
    part := flag.Int("part", 0, "part to execute")
    flag.Parse()

    if *day == 0 {
	fmt.Println("missing day param")
	return
    }

    input_path := fmt.Sprintf("input/day%d.txt", *day)

    f, err := os.Open(input_path)
    if err != nil {
	panic(err)
    }
    defer f.Close()

    var lines []string
    s := bufio.NewScanner(f)
    for s.Scan() {
	lines = append(lines, s.Text())
    }
    
    if err := s.Err(); err != nil {
	panic(err)
    }

    if *part == 0 {
	run(*day, 1, lines)
	run(*day, 2, lines)
    } else {
	run(*day, *part, lines)
    }
}


func run(day, part int, input []string) {
    s := solutions.Solutions{}
    m := reflect.ValueOf(s).MethodByName(fmt.Sprintf("Day%d_%d", day, part))
    m.Call([]reflect.Value{ reflect.ValueOf(input) })
}
