package piscine

func StrRev(str string) string {
	var str2 string
	for _, i := range str {
		str2 += string(i)
	}
	return str2
}
