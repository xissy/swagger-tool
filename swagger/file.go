package swagger

import (
	"io/ioutil"
)

func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteFile(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}
