package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	replacer := strings.NewReplacer(
		"twoneight", "218",
		"eightwone", "821",
		"eighthreeight", "838",
		"sevenineight", "798",
		"oneight", "18",
		"threeight", "38",
		"fiveight", "58",
		"nineight", "98",
		"twone", "21",
		"eightwo", "82",
		"eighthree", "83",
		"sevenine", "79",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	sum := 0

	for scanner.Scan() {
		str := scanner.Text()
		str = replacer.Replace(str)

		regex := regexp.MustCompile("[0-9]+")
		line_numbers := regex.FindAllString(str, -1)

		first, _ := strconv.Atoi(line_numbers[0][0:1])
		last, _ := strconv.Atoi(line_numbers[len(line_numbers)-1])

		line_sum := (first * 10) + (last % 10)
		fmt.Println(line_numbers, " = ", line_sum)
		sum += line_sum
	}

	fmt.Println(sum)
}
