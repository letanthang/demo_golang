package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func main() {
	r1 := strings.NewReader("hello world") //our io.Reader
	r2 := strings.NewReader("hello world")

	hashAndSendNaive(r1)
	hashAndSend(r2)
}

func hashAndSendNaive(r io.Reader) {
	bytes, _ := io.ReadAll(r) //All bytes are now in memory
	fmt.Println(string(bytes))

	w := sha256.New()
	w.Write(bytes)
	sha := hex.EncodeToString(w.Sum(nil))
	fmt.Println(sha)
}

func hashAndSend(r io.Reader) {
	w := sha256.New()

	//any reads from tee will read from r and write to w
	tee := io.TeeReader(r, w)

	sendReader(tee)
	sha := hex.EncodeToString(w.Sum(nil))
	fmt.Println(sha)
}

// sendReader sends the contents of an io.Reader to stdout using a 256 byte buffer
func sendReader(data io.Reader) {
	buff := make([]byte, 256)
	for {
		_, err := data.Read(buff)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buff))
	}
	fmt.Println("")
}
