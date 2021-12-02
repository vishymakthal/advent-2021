package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumWindows(n []int) []int {

	var res []int
	for i := 0; i < len(n)-2; i++ {
		res = append(res, n[i]+n[i+1]+n[i+2])
	}

	return res
}

func countPositiveDiffs(n []int) int {

	var res int
	for i := 1; i < len(n); i++ {
		if n[i]-n[i-1] > 0 {
			res++
		}
	}

	return res
}

func main() {

	f, _ := os.Open(os.Args[1])
	defer func() {
		f.Close()
	}()

	s := bufio.NewScanner(f)
	var n []int
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		n = append(n, i)
	}

	fmt.Printf("Part 1: %d\n", countPositiveDiffs(n))
	fmt.Printf("Part 2: %d\n", countPositiveDiffs(sumWindows(n)))
}
