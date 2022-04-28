package context_cases

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
)

/**
在使用withVaule时要注意四个事项：

1 不建议使用context值传递关键参数，关键参数应该显示的声明出来，不应该隐式处理，context中最好是携带签名、trace_id这类值。
2 因为携带value也是key、value的形式，为了避免context因多个包同时使用context而带来冲突，key建议采用内置类型。
3 上面的例子我们获取trace_id是直接从当前ctx获取的，实际我们也可以获取父context中的value，在获取键值对是，我们先从当前context中查找，没有找到会在从父context中查找该键对应的值直到在某个父context中返回 nil 或者查找到对应的值。
4 context传递的数据中key、value都是interface类型，这种类型编译期无法确定类型，所以不是很安全，所以在类型断言时别忘了保证程序的健壮性。

*/

const (
	KEY1 = "trace_id"
	KEY2 = "event_id"
)

func TestWithValue(t *testing.T) {
	ProcessEnter(NewContextWithTraceID())
}

func NewRequestID() string {
	requestID := strings.Replace(uuid.New().String(), "-", "", -1)
	fmt.Println("request ID = " + requestID)
	return requestID
}

func NewEventID() string {
	requestID := strings.Replace(uuid.New().String(), "-", "", -1)
	fmt.Println("event ID = " + requestID)
	return requestID
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY1, NewRequestID())
	ctx = context.WithValue(ctx, KEY2, NewEventID())
	return ctx
}

func PrintLog(ctx context.Context) {
	fmt.Printf("trace_id=%s\n", GetContextValue(ctx, KEY1))
	fmt.Printf("event_id=%s\n", GetContextValue(ctx, KEY2))
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx)
}
