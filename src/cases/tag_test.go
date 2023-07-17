package cases

import (
	"fmt"
	"reflect"
	"testing"
)

type T1 struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
}

func TestTag(t *testing.T) {
	t1 := reflect.TypeOf(T1{})
	f1, _ := t1.FieldByName("f1")
	fmt.Println(f1.Tag) // f one
	f4, _ := t1.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := t1.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five
}

type T2 struct {
	f string `one:"1" two:"2" blank:""`
}

func TestTag2(t *testing.T) {
	t2 := reflect.TypeOf(T2{})
	f, _ := t2.FieldByName("f")
	fmt.Println(f.Tag) // one:"1" two:"2"blank:""
	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok) // , true
	v, ok = f.Tag.Lookup("five")
	fmt.Printf("%s, %t\n", v, ok) // , false
}

type T3 struct {
	f string "one:`1`"
}

func TestTag3(t *testing.T) {
	t3 := reflect.TypeOf(T3{})
	f, _ := t3.FieldByName("f")
	fmt.Println(f.Tag) // one:`1`
	v, ok := f.Tag.Lookup("one")
	// 即使 tag 是任何字符串值（不管是释义或原始值），只有在双引号（源代码）之间包含值时，Lookup 和 Get 方法才会找到 key 的值
	fmt.Printf("%s, %t\n", v, ok) // , false
}
