package utils

// CheckError is ..
func CheckError(err error) {
	if err != nil {
		err.Error()
	}
}
