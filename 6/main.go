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
	for i, x := range time {
		good_time := 0
		for j := 0; j <= x; j++ {
			sum := j * (x - j)
			if sum >= dist[i] {
				good_time++
			}
		}
		res1 *= good_time
	}

	fmt.Println("Part 1:", res1)

	time2, _ := strconv.Atoi(strings.Join(time_arr, ""))
	dist2, _ := strconv.Atoi(strings.Join(dist_arr, ""))

	res2 := 0
	for i := 0; i <= time2; i++ {
		sum := i * (time2 - i)
		if sum >= dist2 {
			res2++
		}
	}
	fmt.Println("Part 2: ", res2)
}
