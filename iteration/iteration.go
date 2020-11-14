package iteration

// Repeat character input
func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}