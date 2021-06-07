package slices

import (
	"strconv"
	"strings"
)

// JoinInt64 Joins an int64 slice into a string using the passed seperator
func JoinInt64(slice []int64, seperator string) string {
	var strs []string
	for _, number := range slice {
		strs = append(strs, strconv.FormatInt(number, 10))
	}
	return strings.Join(strs, seperator)
}

// SplitInt64 splits a string of int64s by the seperator
func SplitInt64(target, seperator string) []int64 {
	strs := strings.Split(target, seperator)
	var ints = make([]int64, 0, len(strs))
	for _, str := range strs {
		number, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			ints = append(ints, number)
		}
	}
	return ints
}
