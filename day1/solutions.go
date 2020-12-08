package day1

import (
	"fmt"
	"strconv"
)

// Find2020 day1.1
func Find2020(numbers []int) (int, int) {
	var sum int
	start := 0
	end := len(numbers) - 1
	for start < len(numbers)-1 {
		for end > start {
			sum = numbers[start] + numbers[end]
			if sum == 2020 {
				return numbers[start], numbers[end]
			}
			end--
		}
		start++
		end = len(numbers) - 1
	}
	fmt.Println("failure")
	return 0, 0
}

// Find2020_2 day1.2
func Find2020_2(numbers []string) (int, int, int) {
	length := len(numbers)
	fmt.Printf("length: %v ", length)
	for x := 0; x < length-3; x++ {
		for y := 1; y < length-2; y++ {
			for z := 2; z < length-1; z++ {
				a, _ := strconv.Atoi(numbers[x])
				b, _ := strconv.Atoi(numbers[y])
				c, _ := strconv.Atoi(numbers[z])
				//fmt.Printf("%v + %v + %v = %v\n", a, b, c, a+b+c)
				if a+b+c == 2020 {
					return a, b, c
				}
			}
		}
	}

	return 0, 0, 0
}
