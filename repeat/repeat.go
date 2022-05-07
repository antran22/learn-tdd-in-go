package repeat

func Repeat(s string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += s
	}
	return
}
