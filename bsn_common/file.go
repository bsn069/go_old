package bsn_common

import (
	"os"
	"strings"
)

// read file return lines
func ReadFile2Lines(strFilePath, strSep string) ([]string, error) {
	data, err := ReadFile2String(strFilePath)
	if err != nil {
		return nil, err
	}

	strLines := strings.Split(data, strSep)
	return strLines, nil
}

func ReadFile2String(strFilePath string) (string, error) {
	data, err := ReadFile2Byte(strFilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadFile2Byte(strFilePath string) ([]byte, error) {
	file, err := os.Open(strFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var i64FileLength int64
	i64FileLength, err = file.Seek(0, 2)
	if err != nil {
		return nil, err
	}

	file.Seek(0, 0)
	data := make([]byte, i64FileLength)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
