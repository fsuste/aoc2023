package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func allZeroes(row []int) bool {
	for _, el := range row {
		if el != 0 {
			return false
		}
	}
	return true
}

func reverse(row []int) []int {
	res := []int{}
	for i := len(row) - 1; i >= 0; i-- {
		res = append(res, row[i])
	}
	return res
}

func calcNext(row []int) int {
	diff_arr := []int{}
	// fmt.Println(row)
	for i := 0; i < len(row)-1; i++ {
		if row[i] == row[i+1] {
			diff_arr = append(diff_arr, 0)
		} else {
			diff_arr = append(diff_arr, row[i+1]-row[i])
		}
	}
	if allZeroes(diff_arr) {
		return row[len(row)-1]
	}
	return row[len(row)-1] + calcNext(diff_arr)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0
	res2 := 0
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		row := []int{}
		for _, el := range arr {
			n, _ := strconv.Atoi(el)
			row = append(row, n)
		}
		res1 += calcNext(row)
		res2 += calcNext(reverse(row))
	}
	fmt.Println("Part 1: ", res1)
	fmt.Println("Part 2: ", res2)

}
