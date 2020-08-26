package arrayslice

// Sum the slice
func Sum(a []int) int {
	r := 0
	for _, v := range a {
		r += v
	}
	return r
}

func SumAll(a ...[]int) []int {
	var r []int
	for _, v := range a {
		r = append(r, Sum(v))
	}
	return r
}
