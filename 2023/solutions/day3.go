package solutions

import (
    "regexp"
    "strconv"
)

func (Solutions) Day3_1(input []string) int {
    sum := 0
    
    num_pat := regexp.MustCompile("[0-9]+")
    sym_pat := regexp.MustCompile("[-*@%#&=/+$]")
    prev_line := ""

    for idx, line := range input {
	var next_line string
	if idx + 1 < len(input) {
	    next_line = input[idx + 1]
	} else {
	    next_line = ""
	}
	for _, num_idx := range num_pat.FindAllIndex([]byte(line), -1) {
	    start := num_idx[0]
	    end := num_idx[1]
	    num, _ := strconv.ParseInt(num_pat.FindString(line[start:]), 10, 32)

	    sym_idxs := concat(
		sym_pat.FindAllIndex([]byte(prev_line), -1),
		sym_pat.FindAllIndex([]byte(line), -1),
		sym_pat.FindAllIndex([]byte(next_line), -1),
	    )
	    for _, idxs := range sym_idxs {
		sym_idx := idxs[0]	
		if sym_idx >= start - 1 && sym_idx < end + 1 {
		    sum += int(num)
		    break
		}
	    }
	}

	prev_line = line
    }

    return sum
}

func (Solutions) Day3_2(input []string) int {
    sum := 0
    
    sym_pat := regexp.MustCompile("[*]")
    num_pat := regexp.MustCompile("[0-9]+")
    prev_line := ""

    for idx, line := range input {
	gears := []int{}
	stars := 0
	var next_line string
	if idx + 1 < len(input) {
	    next_line = input[idx + 1]
	} else {
	    next_line = ""
	}
	for _, idxs := range sym_pat.FindAllIndex([]byte(line), -1) {
	    stars++
	    sym_idx := idxs[0]
	    
	    nums := find_adjacent_nums(prev_line, sym_idx, num_pat)
	    nums = append(nums, find_adjacent_nums(line, sym_idx, num_pat)...)
	    nums = append(nums, find_adjacent_nums(next_line, sym_idx, num_pat)...)

	    if len(nums) == 2 {
		gears = append(gears, nums[0] * nums[1])
		sum += nums[0] * nums[1]
	    }
	}

	prev_line = line
    }

    return sum
}

func find_adjacent_nums(line string, sym_idx int, num_pat *regexp.Regexp) []int {
    nums := []int{}
    for _, num_idx := range num_pat.FindAllIndex([]byte(line), -1) {
	start := num_idx[0]
	end := num_idx[1]
	if sym_idx >= start - 1 && sym_idx < end + 1 {
	    num, _ := strconv.ParseInt(num_pat.FindString(line[start:]), 10, 32)
	    nums = append(nums, int(num))
	}
    }
    return nums
}

func concat(slices ...[][]int) [][]int {
    var total_len int

    for _, slice := range slices {
        total_len += len(slice)
    }

    arr := make([][]int, 0, total_len)

    for _, slice := range slices {
        arr = append(arr, slice...)
    }

    return arr
}
