package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

func expandUniverse(mat []string) ([]int, []int) {
	x := []int{}
	for i := 0; i < len(mat); i++ {
		s := mat[i]
		if strings.Contains(s, "#") {
			continue
		}
		x = append(x, i)
	}
	y := []int{}
	for i := 0; i < len(mat[0]); i++ {
		str := ""
		for j := 0; j < len(mat); j++ {
			str += string(mat[j][i])
		}
		if strings.Contains(str, "#") {
			continue
		}
		y = append(y, i)

	}
	return x, y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("./input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mat := []string{}
	for scanner.Scan() {
		mat = append(mat, scanner.Text())
	}
	exp_x, exp_y := expandUniverse(mat)

	slices.Sort(exp_x)
	slices.Sort(exp_y)

	points := []Point{}
	for i, line := range mat {
		for j, c := range line {
			if string(c) == "#" {
				points = append(points, Point{i, j})
			}
		}
	}

	exp := 999999
	for i, p := range points {
		for j := len(exp_x) - 1; j >= 0; j-- {
			if p.x > exp_x[j] {
				points[i].x += exp
			}
		}
		for j := len(exp_y) - 1; j >= 0; j-- {
			if p.y > exp_y[j] {
				points[i].y += exp
			}
		}
	}

	sum := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			sum += abs(points[i].x-points[j].x) + abs(points[i].y-points[j].y)
		}
	}
	fmt.Println(sum)
}
