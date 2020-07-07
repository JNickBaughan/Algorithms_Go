package main

import "strings"

func reverse(s string) string {
	var reversed string

	sSlice := strings.Split(s, "")

	for _, char := range sSlice {
		reversed = char + reversed
	}
	return reversed
}
