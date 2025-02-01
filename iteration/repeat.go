package iteration

import "strings"

func Repeat(str string, time int) string {
	var repeated strings.Builder
	for i := 0; i < time; i++ {
		repeated.WriteString(str)
	}
	return repeated.String()
}
