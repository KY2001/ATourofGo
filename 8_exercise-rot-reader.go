package main

import (
	"io"
	"os"
	"strings"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyz"

type rot13Reader struct {
	r io.Reader
}

func (a *rot13Reader) Read(b []byte) (int, error) {
	n, err := a.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			b[i] = 'a' + (b[i]-'a'+13)%26
		} else if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = 'A' + (b[i]-'A'+13)%26
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
