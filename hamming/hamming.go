// Package hamming compares DNA strands and computes mutations
package hamming

import "errors"

const testVersion = 4

// Distance function ensures that strings are the same length, if so the string's DNA are compared for differences. That diff is returned.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Strings are not same length, Hamming can not be computed.")
	} else {
		diff := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff += 1
			}
		}
		return diff, nil
	}
}
