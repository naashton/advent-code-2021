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

func main() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var prev int
	var count int

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if prev == 0 {
			prev = val
		} else if val > prev {
			count += 1
			prev = val
		} else {
			prev = val
		}
	}
	fmt.Println(count)
}
