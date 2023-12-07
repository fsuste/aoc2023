package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const pow = "23456789TJQKA"

type Hand struct {
	cards    []int
	points   int
	strength int
}

func (hand *Hand) calculateHandStrength() {
	dict := make(map[int]int)
	for _, n := range hand.cards {
		dict[n]++
	}
	values := []int{}
	for _, v := range dict {
		values = append(values, v)
	}
	slices.Sort(values)
	// fmt.Println(values)
	if values[len(values)-1] == 3 && values[len(values)-2] == 2 {
		hand.strength = 35
	} else if values[len(values)-1] == 2 && values[len(values)-2] == 2 {
		hand.strength = 25
	} else {
		hand.strength = values[len(values)-1] * 10
	}
}

func compareHands(h1, h2 Hand) int {
	if cmp.Compare(h1.strength, h2.strength) != 0 {
		return cmp.Compare(h1.strength, h2.strength)
	} else {
		for i := 0; i < 5; i++ {
			if cmp.Compare(h1.cards[i], h2.cards[i]) != 0 {
				return cmp.Compare(h1.cards[i], h2.cards[i])
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}

	for scanner.Scan() {
		raw := scanner.Text()
		input := strings.Split(raw, " ")
		hand, rank := input[0], input[1]
		hand_pow := strings.Split(hand, "")
		hand_pow_int := []int{}
		for _, p := range hand_pow {
			hand_pow_int = append(hand_pow_int, bytes.IndexByte([]byte(pow), p[0]))
		}
		points, _ := strconv.Atoi(rank)

		hands = append(hands, Hand{cards: hand_pow_int, points: points})
	}

	for i, _ := range hands {
		hands[i].calculateHandStrength()
	}
	slices.SortFunc(hands, compareHands)

	res := int64(0)
	for i, hand := range hands {
		fmt.Println(hand)
		res += int64((i + 1) * hand.points)
	}

	fmt.Println(res)

}
