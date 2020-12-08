package day5

import (
	"fmt"
	"sort"
)

func nthchar(s string, n int) string {
	r := []rune(s)
	return string(r[n])
}

type node struct {
	left, right *node
	value       string
}

type tree struct {
	root *node
}

func lower(first, last int) (int, int) {
	return first, first + ((last - first) / 2)
}

func upper(first, last int) (int, int) {
	return first + (((last - first) +1 )/ 2), last
}

func getRow(ticket string, first, last, x int) int {
	c := nthchar(ticket, x)
	if x == 6 {
		if c == "F" {
			return first
		}
		return last
	}

	if c == "F" {
		first, last = lower(first, last)
	} else {
		first, last = upper(first, last)
	}
	return getRow(ticket, first, last, x+1)
}

func getSeat(ticket string, first, last, x int) int {
	c := nthchar(ticket, x)
	if x == 9 {
		if c == "L" {
			return first
		}
		return last
	}

	if c == "L" {
		first, last = lower(first, last)
	} else {
		first, last = upper(first, last)
	}
	return getSeat(ticket, first, last, x+1)
}

func findMax(arr []int) int {
	max := 0
	for _, x := range arr {
		if x > max {
			max = x
		}
	}
	return max
}



// Solution2 5.2
func Solution2(data []string) int {
	fmt.Printf("%v rows of data\n", len(data))
	var seats [128][8]int
	ids := make([]int, 0)
	for _, ticket := range data {
		row := getRow(ticket, 0, 127, 0)
		seat := getSeat(ticket, 0, 7, 7)
		id := (row * 8) + seat
		fmt.Printf("Row %v -> %v %v = %v\n", ticket, row, seat, id)
		ids = append(ids, id)
		seats[row][seat] = 1
	}
	sort.Ints(ids)
	for x, id := range(ids) {
		if id+2 == ids[x+1] {
			return id+1
		}
	}
	// return findMax(ids)
	return 0
}

// Solution 5.1
func Solution(data []string) int {
	fmt.Printf("%v rows of data\n", len(data))
	ids := make([]int, 0)
	for _, ticket := range data {
		row := getRow(ticket, 0, 127, 0)
		seat := getSeat(ticket, 0, 7, 7)
		id := (row * 8) + seat
		fmt.Printf("Row %v -> %v %v = %v\n", ticket, row, seat, id)
		ids = append(ids, id)
	}
	//fmt.Println(ids)
	return findMax(ids)

}
