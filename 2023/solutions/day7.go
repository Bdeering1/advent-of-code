package solutions

import (
	"sort"
	"strconv"
	"strings"
)

func (Solutions) Day7_1(input []string) int {
	hands := []Hand{}

	for hand_idx, hand_str := range input {
		toks := strings.Split(hand_str, " ")
		bid, _ := strconv.ParseInt(toks[1], 10, 64)
		hands = append(hands, Hand { [5]int64{}, bid, 0})

		for idx, c := range toks[0] {
			var card_val int64
			switch c {
			case 'A':
				card_val = 14
			case 'K':
				card_val = 13
			case 'Q':
				card_val = 12
			case 'J':
				card_val = 11
			case 'T':
				card_val = 10
			default:
				card_val = int64(c - '0')
			}
			hands[hand_idx].cards[idx] = card_val
		}
		
		card_matches := []int{ 0, 0, 0, 0, 0 }
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 5; j++ {
				if hands[hand_idx].cards[i] == hands[hand_idx].cards[j] {
					card_matches[i]++
				}
			}
		}
		sort.Slice(card_matches, func(i, j int) bool {
			return card_matches[i] > card_matches[j]
		})

		switch card_matches[0] {
		case 4:
			hands[hand_idx].strength = 6 // five of a kind
		case 3:
			hands[hand_idx].strength = 5 // four of a kind
		case 2:
			if card_matches[2] == 1 {
				hands[hand_idx].strength = 4 // full house
			} else {
				hands[hand_idx].strength = 3 // three of a kind
			}
		case 1:
			if card_matches[1] == 1 {
				hands[hand_idx].strength = 2 // two pair
			} else {
				hands[hand_idx].strength = 1 // one pair
			}
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength == hands[j].strength {
			for idx := 0; idx < 5; idx++ {
				if hands[i].cards[idx] != hands[j].cards[idx] {
					return hands[i].cards[idx] < hands[j].cards[idx]
				}
			}
		}
		return hands[i].strength < hands[j].strength
	})

	total_winnings := 0
	for hand_idx, hand := range hands {
		total_winnings += (hand_idx + 1) * int(hand.bid)
	}

	return total_winnings
}

func (Solutions) Day7_2(input []string) int {
	hands := []Hand{}

	for hand_idx, hand_str := range input {
		toks := strings.Split(hand_str, " ")
		bid, _ := strconv.ParseInt(toks[1], 10, 64)
		hands = append(hands, Hand { [5]int64{}, bid, 0})

		for idx, c := range toks[0] {
			var card_val int64
			switch c {
			case 'A':
				card_val = 14
			case 'K':
				card_val = 13
			case 'Q':
				card_val = 12
			case 'J':
				card_val = 1
			case 'T':
				card_val = 10
			default:
				card_val = int64(c - '0')
			}
			hands[hand_idx].cards[idx] = card_val
		}

		card_matches := []int{ 0, 0, 0, 0, 0 }
		for i := 0; i < 4; i++ {
			if hands[hand_idx].cards[i] == 1 {
				continue
			}
			for j := i + 1; j < 5; j++ {
				if hands[hand_idx].cards[j] == 1 {
					continue
				}
				if hands[hand_idx].cards[i] == hands[hand_idx].cards[j] {
					card_matches[i]++
				}
			}
		}
		sort.Slice(card_matches, func(i, j int) bool {
			return card_matches[i] > card_matches[j]
		})

		switch card_matches[0] {
		case 4:
			hands[hand_idx].strength = 6 // five of a kind
		case 3:
			hands[hand_idx].strength = 5 // four of a kind
		case 2:
			if card_matches[2] == 1 {
				hands[hand_idx].strength = 4 // full house
			} else {
				hands[hand_idx].strength = 3 // three of a kind
			}
		case 1:
			if card_matches[1] == 1 {
				hands[hand_idx].strength = 2 // two pair
			} else {
				hands[hand_idx].strength = 1 // one pair
			}
		}

		for _, card := range hands[hand_idx].cards {
			if card == 1 {
				switch hands[hand_idx].strength {
				case 1, 2, 3:
					hands[hand_idx].strength += 2 // one pair, two pair, & three of a kind with joker
				case 6:
					continue // five of a kind (case with 5 jokers)
				default:
					hands[hand_idx].strength++ // high card & four of a kind with joker
				}
			}
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength == hands[j].strength {
			for idx := 0; idx < 5; idx++ {
				if hands[i].cards[idx] != hands[j].cards[idx] {
					return hands[i].cards[idx] < hands[j].cards[idx]
				}
			}
		}
		return hands[i].strength < hands[j].strength
	})

	total_winnings := 0
	for hand_idx, hand := range hands {
		total_winnings += (hand_idx + 1) * int(hand.bid)
	}

	return total_winnings
}

type Hand struct {
	cards [5]int64
	bid int64
	strength int64
}
