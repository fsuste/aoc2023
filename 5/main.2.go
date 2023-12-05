package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Offset struct {
	start, end, diff int64
}

func NewOffset(raw string) Offset {
	sp := strings.Split(raw, " ")
	numbers := []int64{}
	for _, i := range sp {
		n, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}

	return Offset{end: numbers[0], start: numbers[1], diff: numbers[2]}
}

func parseBlock(scanner *bufio.Scanner, seeds []int64) []int64 {
	offsets := []Offset{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		offsets = append(offsets, NewOffset(line))
	}
	for i, s := range seeds {
		for _, o := range offsets {
			if o.start <= s && o.start+o.diff > s {
				seeds[i] += o.end - o.start
				break
			}
		}
	}
	return seeds
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds_text := scanner.Text()
	seeds_text = strings.Split(seeds_text, ": ")[1]
	fmt.Println(seeds_text)
	re := regexp.MustCompile(`[0-9]+ [0-9]+`)
	seeds_pairs := re.FindAllStringSubmatch(seeds_text, -1)
	seeds := []int64{}
	for _, pair := range seeds_pairs {
		pairs := strings.Split(pair[0], " ")
		start, err := strconv.ParseInt(pairs[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		n, err := strconv.ParseInt(pairs[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		for i := int64(0); i < n; i++ {
			seeds = append(seeds, start+i)
		}
	}

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, "map:") {
			seeds = parseBlock(scanner, seeds)
		}

	}
	fmt.Println(slices.Min(seeds))
}
