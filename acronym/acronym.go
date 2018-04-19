// Package acronym creates acronyms from strings
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate creates acronyms from strings
func Abbreviate(s string) string {

	var output string

	// Replace hyphens with spaces
	re := regexp.MustCompile("-")
	s = re.ReplaceAllString(s, " ")

	// Break string into array
	arr := strings.Fields(s)

	// append first letters to output string
	for i := 0; i < len(arr); i++ {
		firstLetter := string(strings.Split(arr[i], "")[0])
		firstLetter = strings.ToUpper(firstLetter)

		output = output + firstLetter
	}

	return output
}
