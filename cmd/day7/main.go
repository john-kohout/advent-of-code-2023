package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"advent-of-code/pkg/file"
)

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

const (
	Ace Card = iota
	King
	Queen
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	Joker
)

type HandType int
type Card int

type Hand struct {
	Bid      int
	Cards    []Card
	Raw      string
	HandType HandType
}

func NewHand(line string) Hand {
	chars := strings.Split(line, " ")[0]
	bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
	var cards []Card
	for _, c := range chars {
		switch c {
		case 'A':
			cards = append(cards, Ace)
		case 'K':
			cards = append(cards, King)
		case 'Q':
			cards = append(cards, Queen)
		case 'J':
			cards = append(cards, Joker)
		case 'T':
			cards = append(cards, Ten)
		case '9':
			cards = append(cards, Nine)
		case '8':
			cards = append(cards, Eight)
		case '7':
			cards = append(cards, Seven)
		case '6':
			cards = append(cards, Six)
		case '5':
			cards = append(cards, Five)
		case '4':
			cards = append(cards, Four)
		case '3':
			cards = append(cards, Three)
		case '2':
			cards = append(cards, Two)
		}
	}

	counter := make(map[Card]int)
	for _, c := range cards {
		if _, ok := counter[c]; ok {
			counter[c]++
			continue
		}
		counter[c] = 1
	}

	var counts []int
	jokers := 0
	for k, count := range counter {
		if k == Joker {
			jokers = count
			continue
		}
		counts = append(counts, count)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if len(counts) == 0 || counts[0] == 5 || (counts[0] == 4 && jokers >= 1) || (counts[0] == 3 && jokers >= 2) || (counts[0] == 2 && jokers >= 3) || jokers >= 4 {
		return Hand{
			Bid:      bid,
			Cards:    cards,
			Raw:      line,
			HandType: FiveOfAKind,
		}
	}
	if counts[0] == 4 || (counts[0] == 3 && jokers == 1) || (counts[0] == 2 && jokers == 2) || (counts[0] == 1 && jokers == 3) {
		return Hand{
			Bid:      bid,
			Cards:    cards,
			Raw:      line,
			HandType: FourOfAKind,
		}
	}
	if counts[0] == 3 || (counts[0] == 2 && jokers == 1) || (counts[0] == 1 && jokers == 2) {
		if counts[1] == 2 {
			return Hand{
				Bid:      bid,
				Cards:    cards,
				Raw:      line,
				HandType: FullHouse,
			}
		} else {
			return Hand{
				Bid:      bid,
				Cards:    cards,
				Raw:      line,
				HandType: ThreeOfAKind,
			}
		}
	}

	if counts[0] == 2 {
		if counts[1] == 2 {
			return Hand{
				Bid:      bid,
				Cards:    cards,
				Raw:      line,
				HandType: TwoPair,
			}
		} else {
			return Hand{
				Bid:      bid,
				Cards:    cards,
				Raw:      line,
				HandType: OnePair,
			}
		}
	}

	if counts[0] == 1 && jokers == 1 {
		return Hand{
			Bid:      bid,
			Cards:    cards,
			Raw:      line,
			HandType: OnePair,
		}
	}

	return Hand{
		Bid:      bid,
		Cards:    cards,
		Raw:      line,
		HandType: HighCard,
	}
}

func main() {
	scanner, f := file.NewScanner("tests/input.txt")
	defer f.Close()
	result := GetResult(scanner)
	fmt.Println(result)
}

func GetResult(scanner *bufio.Scanner) int {
	result := 0

	var hands []Hand
	for scanner.Scan() {
		hands = append(hands, NewHand(scanner.Text()))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].HandType != hands[j].HandType {
			return hands[i].HandType > hands[j].HandType
		}
		for k := 0; k < 5; k++ {
			card1 := hands[i].Cards[k]
			card2 := hands[j].Cards[k]
			if card1 != card2 {
				return card1 > card2
			}
		}
		panic("error")
	})

	for i, hand := range hands {
		result += hand.Bid * (i + 1)
	}

	return result
}
