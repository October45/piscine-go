package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var intstr string = ""
	for _, v := range s {
		if v == '-' && len(intstr) == 0 {
			intstr += "-"
		}
		if v < '0' || v > '9' {
			continue
		}
		intstr += string(v)
	}
	if len(intstr) == 0 || intstr == "-" {
		return 0
	}
	return Atoi(intstr)
}
