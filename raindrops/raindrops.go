// Package raindrops converts numbers to sounds
package raindrops

import (
	"strconv"
)

// Convert takes a number and based on its
func Convert(input int) string {
	output := ""

	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		// apparently string conversions arent normal.
		// use strconv package to convert int to str
		output = strconv.Itoa(input)
	}

	return output
}
