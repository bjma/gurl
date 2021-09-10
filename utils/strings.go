package utils

// Builds a reverse string by truncating towards the middle
func ReverseString(s string) string {
	// see: https://go.dev/blog/strings
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
