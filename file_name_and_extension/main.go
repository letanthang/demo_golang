package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	url := "http://127.0.0.1:5500/203378613949062.pdf"

	fmt.Printf("The url: %s has file_name: %s and extension: %s", url, getFileNameWithoutExt(url), getFileExtension(url))
}

func getFileExtension(filePath string) string {
	ext := filepath.Ext(filePath)
	return ext[1:]
}

func getFileNameWithoutExt(filePath string) string {
	fileName := filepath.Base(filePath)
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
