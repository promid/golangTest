package tdd

import (
	"fmt"
)

type LeaderBoard struct {
	announcer Announce
	wins      map[string]int
}

func (lb *LeaderBoard) win(name string, score int) bool {
	maxer := true
	existingMaxerScore := 0
	existingMaxerName := ""
	for _name, _score := range lb.wins {
		if _score > score {
			maxer = false
		}
		if existingMaxerScore < _score {
			existingMaxerScore = _score
			existingMaxerName = _name
		}
	}
	if lb.wins[name] <= score {
		lb.wins[name] = score
	}
	if maxer && existingMaxerName != name {
		lb.announcer.announceNo1(name + " wins")
	}
	return true
}

type Announce interface {
	announceNo1(string)
}

var _ Announce = &FakeTwitter{}

type FakeTwitter struct {
	anns []string
}

func (f *FakeTwitter) announceNo1(name string) {
	fmt.Println("new winer: ", name)
	f.anns = append(f.anns, name)
}

func (f *FakeTwitter) getAnnouncements() []string {
	return append([]string{}, f.anns...)
}
