package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	elves := []int{}
	count := 0

	for scanner.Scan() {
		currCalories, _ := strconv.Atoi(scanner.Text())

		// Create first item
		if len(elves) == 0 {
			elves = append(elves, currCalories)
			continue
		}

		// Create next item on empty line
		if currCalories == 0 {
			elves = append(elves, currCalories)
			count++
			continue
		}

		elves[count] += currCalories
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Println(elves[0])
	fmt.Println(elves[0] + elves[1] + elves[2])
}
