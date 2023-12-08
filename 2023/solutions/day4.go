package solutions

import (
    "regexp"
    "strings"
)

func (Solutions) Day4_1(input []string) int {
    sum := 0
    
    pat := regexp.MustCompile("[0-9]+")
    col_pat := regexp.MustCompile("[:]")
    for _, line := range input {
	line_score := 0
	substrs := strings.Split(line, "|")
	
	col_idx := col_pat.FindIndex([]byte(line))
	winning_nums := pat.FindAllString(substrs[0][col_idx[0]:], -1)
	my_nums := pat.FindAllString(substrs[1], -1)

	for _, winning_num := range winning_nums {
	    for _, num := range my_nums {
		if num == winning_num {
		    if line_score == 0 {
			line_score = 1
		    } else {
			line_score *= 2
		    }
		}
	    }
	}

	sum += line_score
    }

    return sum
}

func (Solutions) Day4_2(input []string) int {
    total_cards := 0
    
    card_multipliers := [200]int{}
    for i := range card_multipliers {
	card_multipliers[i] = 1
    }

    pat := regexp.MustCompile("[0-9]+")
    col_pat := regexp.MustCompile("[:]")
    for line_idx, line := range input {
	card_score := 0

	substrs := strings.Split(line, "|")
	col_idx := col_pat.FindIndex([]byte(line))
	winning_nums := pat.FindAllString(substrs[0][col_idx[0]:], -1)
	my_nums := pat.FindAllString(substrs[1], -1)

	for _, winning_num := range winning_nums {
	    for _, num := range my_nums {
		if num == winning_num {
		    card_score++
		}
	    }
	}

	for i := 0; i < card_score; i++ {
	    if i + line_idx + 1 > 190 {
		break
	    }
	    card_multipliers[i + line_idx + 1] += card_multipliers[line_idx]
	}

	total_cards += card_multipliers[line_idx]
    }

    return total_cards
}
