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

type Memo struct {
	idx  int
	curr string
	line *Line
}

var mem = make(map[Memo]int)

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

func solve(line Line) int {
	res := 0
	res += backtrack(0, "", line)

	return res
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

func backtrack(idx int, current string, line Line) int {
	//if mem[Memo{idx, current, &line}] != 0 {
	//	fmt.Println("Cache hit")
	//	return mem[Memo{idx, current, &line}]
	//}
	count := 0
	if idx == len(line.s_map) {
		if good(Line{current, line.groups}) {
			// fmt.Println(current)
			count++
		}
		//mem[Memo{idx, current, &line}] = count
		return count
	}
	if string(line.s_map[idx]) == "?" {
		count += backtrack(idx+1, current+"#", line)
		count += backtrack(idx+1, current+".", line)
	}

	return count + backtrack(idx+1, current+string(line.s_map[idx]), line)
}

func good(line Line) bool {
	if strings.Contains(line.s_map, "?") {
		return false
	}
	str_groups := strings.Split(line.s_map, ".")
	str_groups = removeEmpty(str_groups)
	if len(str_groups) != len(line.groups) {
		return false
	}
	viable := true
	for i, n := range line.groups {
		if len(str_groups[i]) != n {
			viable = false
		}
	}
	return viable
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

	res := 0
	for _, l := range mat {
		c := solve(l)
		fmt.Println(l, " - ", c)
		res += c
	}

	fmt.Println(res)
}
