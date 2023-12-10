package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Node struct {
	x, y int
	step int
}

func (n1 Node) add(n2 Node) Node {
	return Node{x: n1.x + n2.x, y: n1.y + n2.y, step: n1.step + n2.step}
}

func (node Node) OutOfBounds(max_x, max_y int) bool {
	if node.x < 0 || node.y < 0 || node.x >= max_x || node.y >= max_y {
		return true
	}
	return false
}

var ValidNodes = [4]Node{Node{-1, 0, 1}, Node{0, -1, 1}, Node{0, 1, 1}, Node{1, 0, 1}}
var ValidCurrent = [4]string{"SL|J", "S-J7", "SLF-", "S|F7"}
var ValidNodeOption = [4]string{"7F|S", "-FLS", "7-JS", "J|LS"}

func BFS(field []string, queue *[]Node) [][]int {
	res_mat := make([][]int, len(field))
	for i := 0; i < len(field); i++ {
		res_mat[i] = make([]int, len(field[0]))
	}
	q := *queue
	for len(q) != 0 {
		el := q[0]
		q = q[1:]
		for i, n := range ValidNodes {
			newNode := el.add(n)
			res_mat[el.x][el.y] = el.step
			//fmt.Println(newNode)
			if newNode.OutOfBounds(len(field), len(field[0])) || res_mat[newNode.x][newNode.y] != 0 {
				continue
			}
			if strings.Contains(ValidCurrent[i], string(field[el.x][el.y])) {
				//fmt.Println(string(field[el.x][el.y]), " -> ", string(field[newNode.x][newNode.y]))

				if strings.Contains(ValidNodeOption[i], string(field[newNode.x][newNode.y])) {
					//fmt.Println("New node in a good path")
					q = append(q, newNode)
				}
			}
		}
	}

	return res_mat
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := []string{}

	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	queue := []Node{}

	start := Node{}
	for i, line := range field {
		for j, el := range line {
			if string(el) == "S" {
				start = Node{i, j, 0}
				queue = append(queue, start)
				break
			}
		}
	}

	res_mat := BFS(field, &queue)

	max := 0
	for _, row := range res_mat {
		row_max := slices.Max(row)
		if max < row_max {
			max = row_max
		}
	}
	// for i, _ := range res_mat {
	// 	fmt.Println(res_mat[i])
	// }

	fmt.Println(max)

	res_mat[start.x][start.y] = 1

	w := len(res_mat[0])
	h := len(res_mat)

	res := 0

	for i, _ := range res_mat {
		for j, _ := range res_mat[i] {
			if res_mat[i][j] != 0 {
				continue
			}

			wall := 0

			x := i
			y := j

			// fmt.Print("(", i, ",", j, ") - ")
			for x < h && y < w {
				if res_mat[x][y] != 0 && string(field[x][y]) != "L" && string(field[x][y]) != "7" {
					wall++
					//fmt.Print("[", x, ",", y, "]")
				}
				x++
				y++
			}
			if wall%2 == 1 {
				res++
				// fmt.Println("(", i, ",", j, ")")
			}
			// fmt.Print("\n")
		}
	}

	// for i, row := range res_mat {
	// 	for j, _ := range row {
	// 		fmt.Printf("%2d ", res_mat[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }

	fmt.Println(res)
}
