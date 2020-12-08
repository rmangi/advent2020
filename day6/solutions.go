package day6

import "fmt"

type survey struct {
	answers [26]int
	responses int
}

func nthchar(s string, n int) rune {
	r := []rune(s)
	return r[n]
}

func scan(results []survey, current *survey, data []string) []survey {
	if len(data) == 0 {
		results = append(results, *current)
		return results
	}
	nextLine := data[0]
	if nextLine == "" {
		results = append(results, *current)
		current = new(survey)
	} else {
		for x := 0; x<len(nextLine);x++ {
			char := nthchar(nextLine, x)
			current.answers[int(char) - 97] = 1
		}
	}
	return scan(results, current, data[1:])
}

func scan2(results []survey, current *survey, data []string) []survey {
	if len(data) == 0 {
		results = append(results, *current)
		return results
	}
	nextLine := data[0]
	if nextLine == "" {
		fmt.Println("survey: ", current)
		results = append(results, *current)
		current = new(survey)
	} else {
		current.responses++
		for x := 0; x<len(nextLine);x++ {
			char := nthchar(nextLine, x)
			current.answers[int(char) - 97]++
		}
	}
	return scan2(results, current, data[1:])
}

// Solution2 6.2
func Solution2(data []string) int {
	fmt.Printf("%v rows of data\n", len(data))
	results := make([]survey, 0)
	current := new(survey)
	answers := scan2(results, current, data)
	fmt.Println(answers)
	total := 0

	for _, a := range(answers) {
		sum := 0 
		for _, y := range(a.answers) {
			if y == a.responses {
				sum++
			}
		}
		total = total + sum
	}

	return total

}


// Solution 6.1
func Solution(data []string) int {
	fmt.Printf("%v rows of data\n", len(data))
	results := make([]survey, 0)
	current := new(survey)
	answers := scan(results, current, data)
	total := 0

	for _, a := range(answers) {
		sum := 0 
		fmt.Println(a)
		for _, y := range(a.answers) {
			if y == 1 {
				sum++
			}
		}
		fmt.Println(sum)
		total = total + sum
	}

	return total

}
