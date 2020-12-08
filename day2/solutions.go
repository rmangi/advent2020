package day2

import (
	"fmt"
	"index/suffixarray"
	"regexp"
	"strconv"
)

// Clue the clue class
type Clue struct {
	min, max int
	val      string
	password string
}

func getintval(s string) int {
	val, err := strconv.Atoi(s)
	if err == nil {
		return val
	}
	fmt.Printf("not an int %v", val)
	return -1
}

func parseData(clues []string) []Clue {
	var listoclues []Clue
	re := regexp.MustCompile(`^(\d+)-(\d+)\s([a-zA-Z]+):\s([a-zA-Z]+)$`)
	for _, c := range clues {
		// fmt.Println(c)
		d := re.FindStringSubmatch(c)
		listoclues = append(listoclues, Clue{getintval(d[1]), getintval(d[2]), d[3], d[4]})
	}
	fmt.Println(len(listoclues), " clues")
	return listoclues
}

func inArray(x int, arr []int) bool {
	for _, i := range arr {
		if i == x {
			return true
		}
	}
	return false
}

// Solution2 solve day2.1
func Solution2(clues []string) int {
	var valid int
	listoclues := parseData(clues)
	for _, c := range listoclues {
		index := suffixarray.New([]byte(c.password))
		matches := index.Lookup([]byte(c.val), -1)
		a := inArray(c.min-1, matches)
		b := inArray(c.max-1, matches)
		if a != b {
			valid++
		}
	}
	return valid
}

// Solution solve day2
func Solution(clues []string) int {
	var valid int
	listoclues := parseData(clues)
	for _, c := range listoclues {
		index := suffixarray.New([]byte(c.password))
		matches := len(index.Lookup([]byte(c.val), -1))
		if matches >= c.min && matches <= c.max {
			valid++
		}
	}
	return valid
}
