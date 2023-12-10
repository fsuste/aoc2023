package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func part1(path string, nav map[string][]string) int {
	res := 0
	curs := "AAA"
	for i := 0; ; i++ {
		if curs == "ZZZ" {
			break
		}
		if string(path[i%len(path)]) == "L" {
			curs = nav[curs][0]
		} else if string(path[i%len(path)]) == "R" {
			curs = nav[curs][1]
		}
		res++
	}
	return res
}

func part2(path string, nav map[string][]string) int {
	start_arr := []string{}
	for k, _ := range nav {
		mat, _ := regexp.MatchString(`..A`, k)
		if mat {
			start_arr = append(start_arr, k)
		}
	}
	round_map := make(map[string][]int)
	res_arr := []int{}
	for _, start := range start_arr {
		i, j := getPhaseAndShift(start, path, nav)
		round_map[start] = []int{i, j}
		res_arr = append(res_arr, j)
	}
	fmt.Println(round_map)

	res := 1
	for _, el := range res_arr {
		res = LCM(res, el)
	}

	return res
}
func GCD(a int, b int) int {
	if a < b {
		return GCD(b, a)
	}
	for a%b != 0 {
		c := a % b
		a = b
		b = c
	}
	return b
}
func LCM(a int, b int) int {
	d := GCD(a, b)
	return (a / d) * b
}

func getElementIndex(slice []string, el any) int {
	for i, x := range slice {
		if x == el {
			return i
		}
	}
	return -1
}

func getPhaseAndShift(start string, path string, nav map[string][]string) (int, int) {
	curs := start
	res_array := []string{curs}
	fmt.Println("Searching path for ", start)
	index := -1
	for i := 0; ; i++ {
		// fmt.Println(curs)
		mat, _ := regexp.MatchString(`..Z`, curs)
		if mat {
			fmt.Println("Z na ", i)
			fmt.Println(i, len(path))
			if string(path[i%len(path)]) == "L" {
				fmt.Println("Printanje lijevo: ", nav[curs][0])
				index = getElementIndex(res_array, nav[curs][0])
			} else if string(path[i%len(path)]) == "R" {
				fmt.Println("Printanje desno: ", nav[curs][1])
				index = getElementIndex(res_array, nav[curs][1])
			}
			break
		}
		if string(path[i%len(path)]) == "L" {
			curs = nav[curs][0]
			res_array = append(res_array, curs)
		} else if string(path[i%len(path)]) == "R" {
			curs = nav[curs][1]
			res_array = append(res_array, curs)
		}
		// fmt.Println(curs)
	}
	// fmt.Println()
	// fmt.Println(strings.Join(res_array, " "))
	// fmt.Println()
	return index, len(res_array) - 1 - index
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path := scanner.Text()

	scanner.Scan()

	nav := make(map[string][]string)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " = ")
		k := parts[0]
		v := strings.Split(strings.Trim(parts[1], "()"), ", ")
		nav[k] = v
	}

	// res1 := part1(path, nav)
	// fmt.Println(res1)
	res2 := part2(path, nav)
	fmt.Println(res2)

}
