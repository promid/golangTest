package tdd

import (
	"testing"
)

func TestTwitter(t *testing.T) {
	ft := &FakeTwitter{}
	var lb = &LeaderBoard{
		announcer: ft,
		wins:      map[string]int{},
	}
	lb.win("tom", 1)
	var annList []string
	annList = ft.getAnnouncements()
	if len(annList) != 1 {
		t.Fatal("expecting 1 announce")
	}

	lb.win("jerry", 3)
	annList = ft.getAnnouncements()
	if len(annList) != 2 {
		t.Fatal("expecting 2 announce")
	}

	lb.win("jack", 1)
	annList = ft.getAnnouncements()
	if len(annList) != 2 {
		t.Fatal("expecting 2 announce")
	}
}
