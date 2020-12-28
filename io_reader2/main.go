package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(bs []byte) (n int, err error) {
	return
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
