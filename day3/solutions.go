package day3

import (
	"fmt"
)

// MatrixMod interface for getting next row/col
type MatrixMod interface {
	nextRow(row int) int
	nextCol(col int) int
}

// MatrixRule rule for modifying each next step
type MatrixRule struct {
	colmod, rmod, maxcol int
}

func nthchar(s string, n int) string {
	r := []rune(s)
	return string(r[n])
}

func (m MatrixRule) nextRow(row int) int {
	return row + m.rmod
}

func (m MatrixRule) nextCol(col int) int {
	return (col + m.colmod) % m.maxcol
}

// Solution2 solution to part 2
func Solution2(trees []string) int {
	r := []rune(trees[0])
	treelength := len(r)
	total := 1

	patterns := []MatrixRule{
		{1, 1, treelength},
		{3, 1, treelength},
		{5, 1, treelength},
		{7, 1, treelength},
		{1, 2, treelength},
	}

	for _, p := range patterns {
		fmt.Println("Next pattern: ", p)
		r := 0
		col := 0
		var treecount int
		for r < len(trees) - 1 {
			r = p.nextRow(r)
			col = p.nextCol(col)
			val := nthchar(trees[r], col)
			if val == "#" {
				treecount++
			}
		}
		fmt.Printf("found %d trees", treecount)
		total = total * treecount
	}

	return total
}

// Solution 3.1
func Solution(trees []string) int {
	fmt.Printf("%v rows of trees to go\n", len(trees))
	var treecount int
	r := []rune(trees[0])
	treelength := len(r)
	spot := 0
	for _, row := range trees[1:] {
		spot += 3
		fmt.Println(row)
		if nthchar(row, spot%treelength) == "#" {
			treecount++
			fmt.Printf("%v is a tree!\n", spot%treelength)
		}
	}
	return treecount
}
