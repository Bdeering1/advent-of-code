package solutions

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
    num_pat := regexp.MustCompile("[0-9]+")

    min_locs := []int64{}
    var wg sync.WaitGroup
    receiver := func(channel <-chan int64) {
	for {
	    val := <- channel
	    fmt.Printf("received %d\n", val)
	    if val == -1 {
		break
	    }
	    min_locs = append(min_locs, val)
	}
    }
    channel := make(chan int64)
    go receiver(channel)

    var seed_start int64 = -1
    for _, seed_str := range num_pat.FindAllString(input[0], -1) {
	num, _ := strconv.ParseInt(seed_str, 10, 64)

	if seed_start == -1 {
	    seed_start = num
	} else {
	    s_start := seed_start
	    seed_range := num
	    channel := channel

	    wg.Add(1)
	    go func() {
		var min_loc int64 = math.MaxInt64
		fmt.Printf("Starting range %d + %d\n", s_start, seed_range)
		defer wg.Done()

		// loop through seed range
		var i int64
		for i = 0; i < seed_range; i++ {
		    map_input := s_start + i
		    modified := false

		    // run input through each map
		    for _, line := range input[2:] {
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
			dest_start, _ := strconv.ParseInt(nums[0], 10, 64)
			input_start, _ := strconv.ParseInt(nums[1], 10, 64)
			range_len, _ := strconv.ParseInt(nums[2], 10, 64)

			// transform input if it falls within input range
			if map_input >= input_start && map_input < input_start + range_len {
			    map_input = dest_start + (map_input - input_start)
			}
		    }
		    // find lowest location #
		    if map_input < min_loc {
			min_loc = map_input
		    }
		}

		channel <- min_loc
	    }()

	    seed_start = -1
	}
    }

    wg.Wait()
    channel <- -1
    close(channel)

    var min_loc int64 =  math.MaxInt64
    for _, loc := range min_locs {
	if loc < min_loc {
	    min_loc  = loc
	}
    }

    return int(min_loc)
}
