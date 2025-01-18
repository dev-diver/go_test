package iteration

func Repeat(str string, time int) string {
	var repeated string
	for i := 0; i < time; i++ {
		repeated = repeated + str
	}
	return repeated
}
