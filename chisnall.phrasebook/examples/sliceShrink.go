package main

func truncate(slice []int) []int {
	var s []int = make([]int, len(slice))
	copy(s, slice)
	return s
}
