package piscine

func ForEach(callback func(int), list []int) {
	for _, v := range list {
		callback(v)
	}
}
