package main

func sum(a []int) int {
	var r int
	for _, val := range a {
		r += val
	}
	return r
}
