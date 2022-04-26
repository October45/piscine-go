package piscine

func Compact(ptr *[]string) int {
	var newList []string
	for _, v := range *ptr {
		if v != "" {
			newList = append(newList, v)
		}
	}
	*ptr = newList
	return len(newList)
}
