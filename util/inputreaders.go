package util


import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadToStrings returns an array of strings from a file
func ReadToStrings(f string) []string {
	var s []string
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		fmt.Println("No file?")
		return s
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		//fmt.Println("got ", val, " ", err)
		s = append(s, val)
		if err != nil {

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return s
}

// ReadToInts read a file to an array of ints
func ReadToInts(f string) []int {
	var numbers []int
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		fmt.Println("No file?")
		return numbers
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		//fmt.Println("got ", val, " ", err)
		numbers = append(numbers, val)
		if err != nil {

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return numbers
}
