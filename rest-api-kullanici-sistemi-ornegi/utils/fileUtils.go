package utils

import (
	"errors"
	"io/ioutil"
)

/*
	Bu helper function'ının amacı gelen dosyayı okumak ve
	elde edilen byte verisini string'e çevirmektir.
*/

// ReadFile is ...
func ReadFile(fileName string) (string, error) {
	if IsEmpty(fileName) {
		return "", errors.New("File is empty")
	}

	bytes, err := ioutil.ReadFile(fileName)
	CheckError(err)
	return string(bytes), nil
}
