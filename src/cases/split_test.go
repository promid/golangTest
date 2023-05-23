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
	t.Log(strings.SplitN(str, "=", 2)) // [a b=c]
	t.Log(strings.SplitN(str, "=", 3)) // [a b c]
	str = "abc"
	t.Log(strings.SplitN(str, "=", 2))      // [abc]
	t.Log(len(strings.SplitN(str, "=", 2))) // 1
	str = "a="
	t.Log(strings.SplitN(str, "=", 2))      // [a ]
	t.Log(len(strings.SplitN(str, "=", 2))) // 2
}

func TestSplit2(t *testing.T) {
	splitN := strings.Split("abc.com", "/")
	t.Log(splitN[0])
	t.Log(splitN[1]) // panic: runtime error: index out of range [1] with length 1
}
