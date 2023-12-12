package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func expandUniverse(mat []string) []string {
	for i := 0; i < len(mat); i++ {
		s := mat[i]
		if strings.Contains(s, "#") {
			continue
		}
		mat = append(mat[:i+1], mat[i:]...)
		newLine := strings.Repeat(".", len(mat[i]))
		mat[i] = newLine
		i++
	}
	ver_pos := []int{}
	for i := 0; i < len(mat[0]); i++ {
		str := ""
		for j := 0; j < len(mat); j++ {
			str += string(mat[j][i])
		}
		if strings.Contains(str, "#") {
			continue
		}
		ver_pos = append(ver_pos, i)

	}
	mat2 := []string{}
	for _, s := range mat {
		str := s
		for i, idx := range ver_pos {
			str = str[:idx+i] + "." + str[idx+i:]
		}
		mat2 = append(mat2, str)

	}
	return mat2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mat := []string{}
	for scanner.Scan() {
		mat = append(mat, scanner.Text())
	}

	mat = expandUniverse(mat)

	points := []Point{}
	for i, line := range mat {
		for j, c := range line {
			if string(c) == "#" {
				points = append(points, Point{i, j})
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
