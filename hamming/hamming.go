// Package hamming calculates hamming distances
package hamming

import (
	"fmt"
)

// Distance takes two strands and returns distance
func Distance(a, b string) (result int, err error) {
	distance := 0

	// throw error if lengths are different
	if len(a) != len(b) {
		return -1, fmt.Errorf("Lengths cannot be different")
	}

	// iterate across both strings at same time,
	// increment if chars don't match
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, err
}
