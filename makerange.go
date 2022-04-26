package piscine

func MakeRange(min int, max int) []int {
	if min >= max {
		return nil
	}
	var array []int
	array = make([]int, max-min)
	for i := 0; i < max-min; i++ {
		array[i] = i + min
	}
	return array
}
