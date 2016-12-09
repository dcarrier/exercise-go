package bob

import (
	"strings"
	//"unicode"
	"fmt"
	"strconv"
)

const testVersion = 2

func Hey(message string) string {
	
	if string(message[len(message)-1]) == "?" {
		message = message[:len(message)-1]
		if message == strings.ToUpper(message) {
			for _, c := range message {
				if string(c) != " " {
					if _, err := strconv.Atoi(string(c)); err != nil {
						if string(c) != ":" {
							fmt.Println(string(c))
							return "Whoa, chill out!"
						}
					}
				}
			}
		}
		return "Sure."
	}
	if message == strings.ToUpper(message) {
		for _, c := range message {
			if string(c) != " " {
				if _, err := strconv.Atoi(string(c)); err != nil {
					if string(c) != "," {
						fmt.Println(string(c))
						return "Whoa, chill out!"
					}
				}
			}
		}
	}
	return "Whatever."
}
