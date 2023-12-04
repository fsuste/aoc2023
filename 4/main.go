package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	winning  []int
	guessing []int
	won      int
}

func (c *Card) calculateWon() {
	slices.Sort(c.winning)
	slices.Sort(c.guessing)
	for _, n := range c.winning {
		for _, g := range c.guessing {
			if n == g {
				c.won++
			}
		}
	}
}

func pow2(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 2
	}
	return res
}

func convToInt(line []string) []int {
	res := []int{}
	for _, i := range line {
		n, _ := strconv.Atoi(i)
		if n != 0 {
			res = append(res, n)
		}
	}
	return res
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := []Card{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		win_part_with_prefix := strings.Split(parts[0], ": ")
		win_part := strings.Split(win_part_with_prefix[1], " ")
		guess_part := strings.Split(parts[1], " ")
		card := Card{winning: convToInt(win_part), guessing: convToInt(guess_part)}

		cards = append(cards, card)
	}

	sum := 0

	for i, _ := range cards {
		cards[i].calculateWon()
		if cards[i].won != 0 {
			sum += pow2(cards[i].won - 1)
		}
	}
	fmt.Println("1st part: ", sum)

	wonMap := make(map[int]int)

	for i, _ := range cards {
		wonMap[i] = 1
	}

	for i, c := range cards {
		fmt.Println("Card number: ", i, " - won: ", c.won)
		for j := i + 1; j <= i+c.won; j++ {
			wonMap[j] += wonMap[i]
		}
	}
	fmt.Println(wonMap)

	sum2 := 0
	for _, v := range wonMap {
		sum2 += v
	}

	fmt.Println(sum2)

}
