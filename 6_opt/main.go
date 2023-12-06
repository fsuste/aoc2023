package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func strArrToInt(str []string) []int {
	res := []int{}
	for _, n := range str {
		t, err := strconv.Atoi(n)
		if err == nil {
			res = append(res, t)
		}
	}
	return res
}

func solveFor(time, dist float64) int {
	d := time*time - 4.0*dist
	x1 := int(math.Floor((time + math.Sqrt(d)) / 2))
	x2 := int(math.Ceil((time - math.Sqrt(d)) / 2))
	return x1 - x2 + 1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\s+`)

	scanner.Scan()
	time_str := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	time_arr := re.Split(time_str, -1)
	time := strArrToInt(time_arr)

	scanner.Scan()
	dist_str := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	dist_arr := re.Split(dist_str, -1)
	dist := strArrToInt(dist_arr)

	res1 := 1
	for i, t := range time {
		res1 *= solveFor(float64(t), float64(dist[i]))
	}
	fmt.Println("Part 1:", res1)

	time2, _ := strconv.Atoi(strings.Join(time_arr, ""))
	dist2, _ := strconv.Atoi(strings.Join(dist_arr, ""))

	res2 := solveFor(float64(time2), float64(dist2))
	fmt.Println("Part 2: ", res2)
}
