package cases

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Integer int

func (a *Integer) Add(b Integer) {
	*a += b
}

func TestAdd(t *testing.T) {
	var a Integer = 1
	a.Add(2)
	if a != 3 {
		t.Error("Add() failed, got", a, "expected 3")
	}
	t.Log("Add() passed")
}

func TestPointerArray(t *testing.T) {
	type Stu struct {
		Name string
		Age  int
	}
	var students []Stu
	students = append(students, Stu{Name: "张三", Age: 18})
	students = append(students, Stu{Name: "李四", Age: 19})

	var studentsP []*Stu

	//slice pointer memory leak
	//wrong ⬇️
	//for _, stu := range students {
	//	studentsP = append(studentsP, &stu)
	//}

	for index := range students {
		studentsP = append(studentsP, &students[index])
	}

	for _, stu := range students {
		stu := stu
		go func() {
			t.Log(stu.Name, stu.Age)
		}()
	}

	time.Sleep(time.Second)
}

type Val struct {
	Int *int
}

func (in *Val) DeepCopy() *Val {
	if in == nil {
		return nil
	}
	out := new(Val)
	in.DeepCopyInto(out)
	return out
}

func (in *Val) DeepCopyInto(out *Val) {
	fmt.Println("out", out)
	fmt.Printf("%p\n", out)
	*out = *in
	out.Int = in.Int
}

func TestVal(t *testing.T) {
	var v *Val = &Val{Int: new(int)}
	copy := v.DeepCopy()
	assert.Equal(t, *v.Int, *copy.Int)

	v = nil
	copy2 := v.DeepCopy()
	assert.Nil(t, copy2)
}
