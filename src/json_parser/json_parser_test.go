package json_parser

import (
	"fmt"
	"testing"

	"github.com/tidwall/gjson"
)

func TestJsonParser(t *testing.T) {
	str := `{"item":"a"}`
	fmt.Println(gjson.Get(str, "item"))
	fmt.Println(gjson.Get(str, ""))
}
