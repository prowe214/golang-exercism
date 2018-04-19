// Package bob defines a conversation with bob
package bob

import (
	"fmt"
	"strings"
)

// Checks if string is all caps
func isAllCaps(remark string) bool {
	return !strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyz")
}

// Hey creates a response
func Hey(remark string) string {
	var answer string

	// input := strings.Trim(remark, "^/0-9!@#$%^&*()_+-=[]{};':\\|,.<>]*$/")
	if remark == "" {
		answer = "Fine. be that way!"
	} else if isAllCaps(remark) && strings.HasSuffix(remark, "?") {
		answer = "Calm down, I know what I'm doing!"
	} else if isAllCaps(remark) {
		answer = "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		answer = "Sure."
	} else {
		answer = "Whatever."
	}

	fmt.Println("----------------")
	fmt.Println(remark)
	fmt.Println(answer)

	return answer
}
