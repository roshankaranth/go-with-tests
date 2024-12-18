package iteration

func Repeat(repeatedCount int, character string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}

	return repeated
}
