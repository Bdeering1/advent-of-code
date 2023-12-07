package solutions

import (
    "regexp"
    "strconv"
)

func (Solutions) Day2_1(input []string) int {
    sum := 0
    red_pat := regexp.MustCompile("([0-9])+ r")
    green_pat := regexp.MustCompile("([0-9])+ g")
    blue_pat := regexp.MustCompile("([0-9])+ b")

    LineLoop:
    for idx, line := range input {
	for _, val := range red_pat.FindAllString(line, -1) {
	    num, _ := strconv.ParseInt(val[:len(val) - 2], 10, 32)
	    if num > 12 {
		continue LineLoop
	    }
	}
	for _, val := range green_pat.FindAllString(line, -1) {
	    num, _ := strconv.ParseInt(val[:len(val) - 2], 10, 32)
	    if num > 13 {
		continue LineLoop
	    }
	}
	for _, val := range blue_pat.FindAllString(line, -1) {
	    num, _ := strconv.ParseInt(val[:len(val) - 2], 10, 32)
	    if num > 14 {
		continue LineLoop
	    }
	}
	sum += idx + 1
    }

    return sum
}

func (Solutions) Day2_2(input []string) int {
    sum := 0
    red_pat := regexp.MustCompile("([0-9])+ r")
    green_pat := regexp.MustCompile("([0-9])+ g")
    blue_pat := regexp.MustCompile("([0-9])+ b")

    for _, line := range input {
	var max_red int64 = 0
	for _, val := range red_pat.FindAllString(line, -1) {
	    num, _ := strconv.ParseInt(val[:len(val) - 2], 10, 32)
	    if num > max_red {
		max_red = num
	    }
	}
	var max_green int64 = 0
	for _, val := range green_pat.FindAllString(line, -1) {
	    num, _ := strconv.ParseInt(val[:len(val) - 2], 10, 32)
	    if num > max_green {
		max_green = num
	    }
	}
	var max_blue int64 = 0
	for _, val := range blue_pat.FindAllString(line, -1) {
	    num, _ := strconv.ParseInt(val[:len(val) - 2], 10, 32)
	    if num > max_blue {
		max_blue = num
	    }
	}

	sum += int(max_red * max_green * max_blue)
    }

    return sum
}
