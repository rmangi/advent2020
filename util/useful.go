package util

// Nthchar - read nth char from a string as a list of runes
func Nthchar(s string, n int) string {
	r := []rune(s)
	return string(r[n])
}

