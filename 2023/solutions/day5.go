package solutions

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func (Solutions) Day5_1(input []string) int {
	num_pat := regexp.MustCompile("[0-9]+")

	// gather seed numbers
	inputs := []int64{}
	modified := []bool{}
	for _, seed_str := range num_pat.FindAllString(input[0], -1) {
		seed, _ := strconv.ParseInt(seed_str, 10, 64)
		inputs = append(inputs, seed)
		modified = append(modified, false)
	}

	// for each map, compute corresponding values for each of the inputs
	// ie. for seed/soil map { list_of_inputs -> list_of_soils }
	for _, line := range input[2:] {
		if strings.Contains(line, ":") {
			// start map
			for i := range inputs {
				modified[i] = false
			}
			continue
		}
		if len(line) == 0 {
			// end map
			continue
		}

		nums := num_pat.FindAllString(line, 3)
		dest_start, _ := strconv.ParseInt(nums[0], 10, 64)
		input_start, _ := strconv.ParseInt(nums[1], 10, 64)
		range_len, _ := strconv.ParseInt(nums[2], 10, 64)
		for idx, input := range inputs {
			// transform input if it falls within input range
			if !modified[idx] && input >= input_start && input < input_start+range_len {
				inputs[idx] = dest_start + (input - input_start)
				modified[idx] = true
			}
		}
	}

	// find lowest location #
	var min_loc int64 = math.MaxInt64
	for _, input := range inputs {
		if input < min_loc {
			min_loc = input
		}
	}

	return int(min_loc)
}

func (Solutions) Day5_2(input []string) int {
	num_pat := regexp.MustCompile("[0-9]+")
	num_pair := regexp.MustCompile("[0-9]+ [0-9]+")

	// gather seed + location ranges
	seed_ranges := [][]int64{}
	for _, seed_str := range num_pair.FindAllString(input[0], -1) {
		toks := strings.Split(seed_str, " ")
		start, _ := strconv.ParseInt(toks[0], 10, 64)
		range_len, _ := strconv.ParseInt(toks[1], 10, 64)
		seed_ranges = append(seed_ranges, []int64{start, range_len})
	}
	loc_ranges := [][]int64{}
	for i := len(input) - 2; i >= 0; i-- {
		if strings.Contains(input[i], ":") {
			break
		}
		nums := num_pat.FindAllString(input[i], 3)
		start, _ := strconv.ParseInt(nums[0], 10, 64)
		range_len, _ := strconv.ParseInt(nums[2], 10, 64)
		loc_ranges = append(loc_ranges, []int64{start, range_len})
	}

	// starting from the smallest location strings, transform backwards through each map until a seed is found
	sort.Slice(loc_ranges, func(i, j int) bool { return loc_ranges[i][0] < loc_ranges[j][0] })
	for _, loc_range := range loc_ranges {
		loc_start := loc_range[0]

		var i int64
		for i = 0; i < loc_range[1]; i++ {
			map_input := loc_start + i
			modified := false

			// run input through each map
			for i := len(input) - 2; i > 0; i-- {
				line := input[i]

				if strings.Contains(line, ":") {
					// start map
					modified = false
					continue
				}
				if modified || len(line) == 0 {
					// end map
					continue
				}

				nums := num_pat.FindAllString(line, 3)
				input_start, _ := strconv.ParseInt(nums[0], 10, 32)
				dest_start, _ := strconv.ParseInt(nums[1], 10, 32)
				range_len, _ := strconv.ParseInt(nums[2], 10, 32)

				// transform input if it falls within input range
				if map_input >= input_start && map_input < input_start+range_len {
					map_input = dest_start + (map_input - input_start)
					modified = true
				}
			}

			// check if there is a matching seed
			for _, seed_range := range seed_ranges {
				if map_input >= seed_range[0] && map_input < seed_range[0]+seed_range[1] {
					return int(loc_start)
				}
			}
		}
	}

	return -1
}
