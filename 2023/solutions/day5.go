package solutions

import (
    "math"
    "regexp"
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
	    if !modified[idx] && input >= input_start && input < input_start + range_len {
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
    return -1
}
