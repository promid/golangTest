package wait_cases

import (
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

func TestWaitGroup(t *testing.T) {
	group := wait.Group{}
	group.Start(func() {
		klog.Infof("work 1 starts")
		time.Sleep(3 * time.Second)
		klog.Infof("work 1 ends")
	})
	group.Start(func() {
		klog.Infof("work 2 starts")
		time.Sleep(2 * time.Second)
		klog.Infof("work 2 ends")
	})
	group.Wait()
}
