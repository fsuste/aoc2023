package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		str := scanner.Text()
		regex := regexp.MustCompile("[0-9]+")
		line_numbers := regex.FindAllString(str, -1)
		// fmt.Println(reflect.TypeOf(line_numbers[0]))
		first, _ := strconv.Atoi(line_numbers[0][0:1])
		last, _ := strconv.Atoi(line_numbers[len(line_numbers)-1])

		line_sum := (first * 10) + (last % 10)
		sum += line_sum
	}

	fmt.Println(sum)
}
