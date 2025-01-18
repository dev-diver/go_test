package iteration

func Repeat(str string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + str
	}
	return repeated
}
