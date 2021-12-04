package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func count_depth_increments(depths []int) {
	var prev int
	var count int

	for _, depth := range depths {
		if prev == 0 {
			prev = depth
		} else if depth > prev {
			count += 1
			prev = depth
		} else {
			prev = depth
		}
	}

	fmt.Println("Number of depth increments is: ", count)
}

func count_summed_depth_increments(depths []int, window_size int) {
	var count int

	for idx, _ := range depths {
		if sum_arr(depths[idx+1:idx+window_size+1]) > sum_arr(depths[idx:idx+window_size]) {
			count += 1
		}
	}
	fmt.Println("Number of depth increments for window size {", window_size, "} is: ", count)
}

func sum_arr(depths []int) int{
	depth_total := 0
	for _, depth := range depths {
		depth_total += depth
	}
	return depth_total
}

func main() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var depths []int

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		depths = append(depths, val)
	}

	count_depth_increments(depths)
	count_summed_depth_increments(depths, 3)
}
