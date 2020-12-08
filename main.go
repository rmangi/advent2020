package main

import (
	"fmt"
	"sort"
	"rmangi.com/user/advent/util"
	"rmangi.com/user/advent/day1"
	"rmangi.com/user/advent/day2"
	"rmangi.com/user/advent/day3"
	"rmangi.com/user/advent/day4"
	"rmangi.com/user/advent/day5"
	"rmangi.com/user/advent/day6"

)

func runday6() {
	data := util.ReadToStrings("data/day6.txt")
	fmt.Println("answer: ", day6.Solution2(data))
}


func runday5() {
	data := util.ReadToStrings("data/day5.txt")
	fmt.Println("answer: ", day5.Solution2(data))
}


func runday4() {
	data := util.ReadToStrings("data/day4.txt")
	fmt.Println("answer: ", day4.Solution2(data))
}

func runday3() {
	clues := util.ReadToStrings("data/day3.txt")
	//fmt.Println(day3.Solution(clues))
	fmt.Println("part 2")
	
	fmt.Println("answer: ", day3.Solution2(clues))
}

func runday2() {
	clues := util.ReadToStrings("data/day2.txt")
	fmt.Println(day2.Solution(clues))
	fmt.Println(day2.Solution2(clues))
}

func runday1() {
	numbers := util.ReadToInts("data/day1.txt")
	sort.Ints(numbers)
	a, b := day1.Find2020(numbers)
	fmt.Println(a, " ", b, " -> ", a*b)

	strings := util.ReadToStrings("day1.txt")
	fmt.Println(strings)
	x, y, z := day1.Find2020_2(strings)
	fmt.Printf("%v, %v, %v -> %v", x, y, z, x*y*z)
}

func main() {
	fmt.Println("Hello, world.")
	runday6()
}
