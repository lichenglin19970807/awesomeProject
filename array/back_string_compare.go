package array_study

func SubCompare(s string) string {
	b := []byte(s)
	slow := 0
	for _, v := range b {
		if v != '#' {
			b[slow] = v
			slow++
		} else if slow != 0 {
			slow--
		}
	}
	return string(b[:slow])
}
