package cases

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func fill() (a_cool_map map[string]string) {
	// This fixes it: a_cool_map = make(map[string]string)
	a_cool_map["key"] = "value"
	return
}

func TestMap(t *testing.T) {
	a_cool_map := fill()
	fmt.Println(a_cool_map)
}

func TestMapDelete(t *testing.T) {
	m := map[string]string{
		"a": "a",
	}
	// deleting non-existing key is OK
	delete(m, "b")
	println(m)
	// reading non-existing key returns empty
	println(m["abcd"])
}

func TestReplaceMapItem(t *testing.T) {
	m := map[string]string{"a": "a"}
	t.Log(m)
	replaceMapItem(m) // it works
	t.Log(m)
}

func replaceMapItem(m map[string]string) {
	m["a"] = "b"
}

func TestReadNilMap(t *testing.T) {
	var m map[string]string
	if _, ok := m["test"]; ok {
		fmt.Println("test")
	}
}

func TestSyncMap(t *testing.T) {
	a := sync.Map{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func(ctx context.Context, m *sync.Map) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				m.Range(func(key, value any) bool {
					fmt.Println("read key value from sync map", key, value)
					return true
				})
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx, &a)
	num := 0
	b := false
	for {
		if b {
			break
		}
		select {
		case <-ctx.Done():
			b = true
		default:
			a.Store(fmt.Sprintf("key-%d", num), fmt.Sprintf("value-%d", num))
			num++
			time.Sleep(300 * time.Millisecond)
		}
	}
	<-ctx.Done()
	fmt.Println("end")
}
