package cases

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	str := "abc:"
	t.Log(len(strings.Split(str, ":")))
	str = "abc"
	t.Log(len(strings.Split(str, ":")))
	str = "kind-"
	t.Log(len(strings.Split(str, "kind-")))
}
