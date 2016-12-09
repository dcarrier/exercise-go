package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

// Convert tests if a factorial of a given array is present, appends a string buffer, and returns the value
func Convert(num int) string {

	var buffer bytes.Buffer
	a := [3]int{3, 5, 7}
	counter := 0

	for i := 0; i <= 2; i++ {
		if num%a[i] == 0 {
			switch a[i] {
			case 3:
				buffer.WriteString("Pling")
			case 5:
				buffer.WriteString("Plang")
			case 7:
				buffer.WriteString("Plong")
			}
		} else {
			counter += 1
			if counter == len(a) {
				return strconv.Itoa(num)
			}
		}
	}
	return buffer.String()
}
