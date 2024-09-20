package iteration

func Repeat(text string, repeatedCount int) string {

	var result string
	for i := 0; i < repeatedCount; i++ {
		result += text
	}
	return result
}
