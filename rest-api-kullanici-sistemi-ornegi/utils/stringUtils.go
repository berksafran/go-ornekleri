package utils

// IsEmpty Dosyanın boş olup olmadığına bakar.
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}
