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

type Round struct {
	line             string
	green, red, blue int
	valid            bool
}

type Game struct {
	line   string
	rounds []*Round
	number int
	valid  bool
}

func (g *Game) parse_game() {
	game_re := regexp.MustCompile(`Game ([0-9]+): `)
	game_match := game_re.FindStringSubmatch(g.line)[1]
	g.number, _ = strconv.Atoi(game_match)

	_, round_lines, _ := strings.Cut(g.line, ": ")
	// fmt.Println(round_lines)
	round_lines_parsed := strings.Split(round_lines, "; ")
	// fmt.Println(len(round_lines_parsed))
	for _, rl := range round_lines_parsed {
		g.rounds = append(g.rounds, &Round{line: rl, valid: true})
	}
	// fmt.Println(g.rounds)

	for _, r := range g.rounds {
		r.parse_round()
		// fmt.Println(r)
		if r.valid == false {
			g.valid = false
		}
	}

}

func (r *Round) parse_round() {
	//fmt.Println(r.line)
	green_rx := regexp.MustCompile(`([0-9]*) green`)
	green := green_rx.FindStringSubmatch(r.line)
	if green != nil {
		r.green, _ = strconv.Atoi(green[1])
		if r.green > 13 {
			r.valid = false
		}
	}
	//fmt.Println("Green: ", r.green)

	red_rx := regexp.MustCompile(`([0-9]*) red`)
	red := red_rx.FindStringSubmatch(r.line)
	if red != nil {
		r.red, _ = strconv.Atoi(red[1])
		if r.red > 12 {
			r.valid = false
		}
	}
	// fmt.Println("Red: ", r.red)

	blue_rx := regexp.MustCompile(`([0-9]*) blue`)
	blue := blue_rx.FindStringSubmatch(r.line)
	if blue != nil {
		r.blue, _ = strconv.Atoi(blue[1])
		if r.blue > 14 {
			r.valid = false
		}
	}
	//fmt.Println("Blue: ", r.blue)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valid_sum := 0
	cube_power_sum := 0

	for scanner.Scan() {
		str := scanner.Text()
		game := Game{line: str, valid: true}
		game.parse_game()
		if game.valid {
			valid_sum += game.number
		}
		r, g, b := 0, 0, 0
		cube_power := 0
		for _, ro := range game.rounds {
			//fmt.Println(ro)
			if r < ro.red {
				r = ro.red
			}
			if g < ro.green {
				g = ro.green
			}
			if b < ro.blue {
				b = ro.blue
			}
		}
		cube_power = r * b * g
		cube_power_sum += cube_power
	}

	fmt.Println("Valid Sum: ", valid_sum)
	fmt.Println("Cube Power Sum", cube_power_sum)

}
