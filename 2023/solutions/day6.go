package solutions

import (
	"regexp"
	"strconv"
)

func (Solutions) Day6_1(input []string) int {
	num_pat := regexp.MustCompile("[0-9]+")

	times := num_pat.FindAllString(input[0], -1)
	dists := num_pat.FindAllString(input[1], -1)

	product := 1
	for race := 0; race < len(times); race++ {
		race_time, _ := strconv.ParseInt(times[race], 10, 32)
		dist_record, _ := strconv.ParseInt(dists[race], 10, 32)
		ways_to_win := race_time

		var i int64
		for i = 0; i < race_time; i++ {
			if i*(race_time-i) <= dist_record {
				ways_to_win--
			}
		}
		product *= int(ways_to_win)
	}

	return product
}

func (Solutions) Day6_2(input []string) int {
	num_pat := regexp.MustCompile("[0-9]+")

	time_nums := num_pat.FindAllString(input[0], -1)
	dist_nums := num_pat.FindAllString(input[1], -1)

	time_str := ""
	for _, num_str := range time_nums {
		time_str += num_str
	}
	dist_str := ""
	for _, num_str := range dist_nums {
		dist_str += num_str
	}

	race_time, _ := strconv.ParseInt(time_str, 10, 64)
	dist_record, _ := strconv.ParseInt(dist_str, 10, 64)
	ways_to_win := race_time

	var i int64
	for i = 0; i < race_time; i++ {
		if i*(race_time-i) <= dist_record {
			ways_to_win--
		}
	}

	return int(ways_to_win)
}
