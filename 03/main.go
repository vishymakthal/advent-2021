package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func binaryToInt(binary string) int {
	res, _ := strconv.ParseInt(binary, 2, 64)
	return int(res)
}

func findCounts(report []string) []int {

	numBits := len(report[0])
	var counts = make([]int, numBits)

	for _, n := range report {
		for i := 0; i < numBits; i++ {
			bit, _ := strconv.Atoi(string(n[i]))
			counts[i] += bit * 1
		}
	}

	return counts
}

func findPowerConsumption(report []string) int {

	counts := findCounts(report)

	var gamma string
	var epsilon string
	for _, n := range counts {
		if n < (len(report) / 2) {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return binaryToInt(gamma) * binaryToInt(epsilon)
}

func findLifeSupportRating(report []string) int {

	numBits := len(report[0])
	var set []string
	set = report
	for i := 0; i < numBits; i++ {
		counts := findCounts(set)
		newSet := []string{}
		ones := float64(counts[i]) >= (float64(len(set)) / 2.0)
		for _, n := range set {
			if ones && n[i] == '1' {
				newSet = append(newSet, n)
			} else if !ones && n[i] == '0' {
				newSet = append(newSet, n)
			}
		}
		set = newSet
		if len(set) == 1 {
			break
		}
	}

	oxygen := binaryToInt(set[0])

	set = report
	for i := 0; i < numBits; i++ {
		counts := findCounts(set)
		newSet := []string{}
		ones := float64(counts[i]) >= (float64(len(set)) / 2.0)
		for _, n := range set {
			if !ones && n[i] == '1' {
				newSet = append(newSet, n)
			} else if ones && n[i] == '0' {
				newSet = append(newSet, n)
			}
		}
		set = newSet
		if len(set) == 1 {
			break
		}
	}

	c02 := binaryToInt(set[0])

	return oxygen * c02
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
	fmt.Printf("Part 1: %d\n", findPowerConsumption((readInput(os.Args[1]))))
	fmt.Printf("Part 2: %d\n", findLifeSupportRating((readInput(os.Args[1]))))
}
