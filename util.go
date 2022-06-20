package main

import "os"

func openOrCreateMru(filePath string) (*os.File, error) {
	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		f, err = os.Create(filePath)
	}
	return f, err
}
