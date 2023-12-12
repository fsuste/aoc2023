package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	s_map  string
	groups []int
}

var memo [200][200][200]int

func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseLine(str string) Line {
	parts := strings.Split(str, " ")
	groups := []int{}
	for _, n := range strings.Split(parts[1], ",") {
		groups = append(groups, toInt(n))
	}
	return Line{parts[0], groups}
}
func removeEmpty(str []string) []string {
	var res []string
	for _, s := range str {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}

func solve(line Line) (int, int) {
	res1 := solveLine(1, line.s_map, line.groups)
	res2 := solveLine(5, line.s_map, line.groups)

	return res1, res2
}

func solveLine(n int, str string, nums []int) int {
	pattern := strings.Builder{}
	numbers := []int{}
	for i := 0; i < n; i++ {
		pattern.WriteString(str)
		if i != n-1 {
			pattern.WriteRune('?')
		}
		numbers = append(numbers, nums...)
	}

	memo = [200][200][200]int{}
	s := pattern.String()
	return backtrack(s, 0, numbers, 0, 0)
}

func backtrack(str string, sidx int, numbers []int, nidx int, group_size int) int {
	if len(str) == sidx { // if on the end of the string
		// fmt.Println("End of string")
		if nidx == len(numbers)-1 && numbers[nidx] == group_size { // valid string
			//fmt.Println("1. Valid, adding 1")
			return 1
		}
		if nidx == len(numbers) && group_size == 0 { // valid string
			//fmt.Println("2. Valid, adding 1")
			return 1
		}
		// fmt.Println("Invalid string :", str, " - ", numbers)
		return 0 // invalid string
	}
	// fmt.Println("Passing to calculate: ", str, " - ", numbers)

	if memo[sidx][nidx][group_size] != 0 {
		return memo[sidx][nidx][group_size] - 1 // -1 for special case of last group
	}

	sum := 0

	if string(str[sidx]) == "?" || string(str[sidx]) == "#" { // opens or continues a group
		sum += backtrack(str, sidx+1, numbers, nidx, group_size+1)
	}
	if string(str[sidx]) == "?" || string(str[sidx]) == "." { //check if I can close a group
		if group_size > 0 && nidx < len(numbers) && numbers[nidx] == group_size { // group exists, group right size and not last
			sum += backtrack(str, sidx+1, numbers, nidx+1, 0)
		}
		if group_size == 0 {
			sum += backtrack(str, sidx+1, numbers, nidx, 0)
		}
	}

	memo[sidx][nidx][group_size] = sum + 1
	return sum
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mat := []Line{}
	for scanner.Scan() {
		mat = append(mat, parseLine(scanner.Text()))
	}

	res1 := 0
	res2 := 0
	for _, l := range mat {
		s1, s2 := solve(l)
		res1 += s1
		res2 += s2
	}

	fmt.Println("Part 1: ", res1)
	fmt.Println("Part 2: ", res2)
}
