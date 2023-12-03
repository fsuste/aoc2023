package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var adjacent = []int{-1, 0, 1}
var lines []string

func findNearbyNumbers(x, y int) int {
	sum := 0
	last := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if unicode.IsDigit(rune(lines[x+adjacent[i]][y+adjacent[j]])) {
				temp := composeDigit(x+adjacent[i], y+adjacent[j])
				if temp != last {
					sum += temp
					last = temp
				}
			}
		}
	}
	return sum
}

func composeDigit(x, y int) int {
	// fmt.Println("Composing number at: ", x, y)
	number := 0
	if y == 0 {
		re := regexp.MustCompile(`[0-9]+`)
		number, _ = strconv.Atoi(re.FindStringSubmatch(lines[x])[0])
		//fmt.Println("Found ", number)
	} else if unicode.IsDigit(rune(lines[x][y])) {
		number = composeDigit(x, y-1)
	} else {
		re := regexp.MustCompile(`[0-9]+`)
		// fmt.Println(re.FindStringSubmatch(lines[x][y-1:])[0])
		number, _ = strconv.Atoi(re.FindStringSubmatch(lines[x][y:])[0])
		// fmt.Println("Found ", number, " at positions: ", x, " ", y)
	}
	// fmt.Println(number)
	return number
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0

	for i, line := range lines {
		for j, char := range line {
			if !unicode.IsDigit(char) && string(char) != "." {
				//fmt.Prtf("Found %c\n", char)
				sum += findNearbyNumbers(i, j)
			}
			// fmt.Printf("%c", char)
		}
	}
	fmt.Println(sum)

}
