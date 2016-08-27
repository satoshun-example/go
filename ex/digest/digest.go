package main

import "fmt"

const ldigsts = "0123456789abcdef"

func digest(b []byte) string {
	buf := make([]byte, 0, len(b)*2)
	for i := range b {
		c := b[i]
		buf = append(buf, ldigsts[c>>4], ldigsts[c&0x0F])
	}

	return string(buf)
}

func main() {
	b := []byte("abcP@-01m>,@0000")
	fmt.Printf("%x\n", b)
	fmt.Println(digest(b))

	fmt.Println(digest(b) == fmt.Sprintf("%x", b))
}
