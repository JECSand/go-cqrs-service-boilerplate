package utilities

import "strconv"

// StringToInt converts a string into an int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
