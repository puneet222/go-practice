// Package word is utility for any given string to analyse its words
package word

import "strings"

// UseCount function count all the occurrences of words
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count function takes the string and returns
// the count of total words of the string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}