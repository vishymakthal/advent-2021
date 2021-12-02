package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	F = "forward"
	D = "down"
	U = "up"
)

func getFinalPosition(inst []string) int {

	var x = 0
	var y = 0

	for _, ln := range inst {
		split := strings.Split(ln, " ")
		units, _ := strconv.Atoi(split[1])

		switch split[0] {
		case F:
			x += units
		case D:
			y += units
		case U:
			y -= units
		}
	}

	return x * y
}

func getFinalPositionWithAim(inst []string) int {

	var x = 0
	var y = 0
	var aim = 0

	for _, ln := range inst {
		split := strings.Split(ln, " ")
		units, _ := strconv.Atoi(split[1])

		switch split[0] {
		case F:
			x += units
			y += units * aim
		case D:
			aim += units
		case U:
			aim -= units
		}
	}

	return x * y
}

func readInput(fname string) []string {

	f, _ := os.Open(fname)
	defer func() {
		f.Close()
	}()

	s := bufio.NewScanner(f)
	var n []string
	for s.Scan() {
		n = append(n, s.Text())
	}

	return n
}

func main() {

	fmt.Printf("Part 1: %d\n", getFinalPosition(readInput(os.Args[1])))
	fmt.Printf("Part 2: %d\n", getFinalPositionWithAim(readInput(os.Args[1])))
}
