package filex

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ToStr(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("fail to read [%s], err: [%v]", filePath, err)
	}
	return string(b), nil
}

func ToBytes(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("fail to read [%s], err: [%v]", filePath, err)
	}
	return b, nil
}

func ToTrimStr(filePath string) (string, error) {
	str, err := ToStr(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
