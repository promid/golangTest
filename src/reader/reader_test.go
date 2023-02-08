package reader

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("abcd\napiVersion: v1\n456:788\nabc:ncc\n"))
	c := "\n"
	readNext := false
	for {
		line, err := r.ReadString('\n')
		if line == "" && err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		if strings.HasPrefix(line, "apiVersion:") || readNext {
			c += line
			readNext = true
		}
	}
	t.Log(c)
}
