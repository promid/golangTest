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
	str = ""
	t.Log(len(strings.Split(str, ":")))
}

func TestSplitN(t *testing.T) {
	str := "a=b=c"
	t.Log(strings.SplitN(str, "=", 2))
	t.Log(strings.SplitN(str, "=", 3))
	str = "abc"
	t.Log(strings.SplitN(str, "=", 2))
	t.Log(len(strings.SplitN(str, "=", 2)))
	str = "a="
	t.Log(strings.SplitN(str, "=", 2))
	t.Log(len(strings.SplitN(str, "=", 2)))
}
