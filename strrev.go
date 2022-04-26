package piscine

func StrRev(str string) string {
	var str2 string
	for i := range str {
		str2 += string(str[len(str)-i-1])
	}
	return str2
}
