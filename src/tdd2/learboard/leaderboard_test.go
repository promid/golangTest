package learboard

import (
	"testing"
)

// scenario A
func TestA(t *testing.T) {
	frank := &person{
		name:  "Frank",
		score: 100,
	}
	peter := &person{
		name:  "Peter",
		score: 200,
	}
	leaderboard := &leaderboard{}
	leaderboard.add(frank)
	leaderboard.add(peter)
	frankRank := leaderboard.getPersonRank(frank)
	if frankRank != 2 {
		t.Errorf("expected 2, actual %d", frankRank)
		return
	}
	peterRank := leaderboard.getPersonRank(peter)
	if peterRank != 1 {
		t.Errorf("expected 1, actual %d", peterRank)
		return
	}
}

// scenario B
func TestB(t *testing.T) {
	frank := &person{
		name:  "Frank",
		score: 250,
	}
	peter := &person{
		name:  "Peter",
		score: 200,
	}
	mary := &person{
		name:  "Peter",
		score: 250,
	}
	leaderboard := &leaderboard{}
	leaderboard.add(frank)
	leaderboard.add(peter)
	leaderboard.add(mary)
	frankRank := leaderboard.getPersonRank(frank)
	if frankRank != 1 {
		t.Errorf("expected 1, actual %d", frankRank)
		return
	}
	maryRank := leaderboard.getPersonRank(mary)
	if maryRank != 1 {
		t.Errorf("expected 1, actual %d", frankRank)
		return
	}
	peterRank := leaderboard.getPersonRank(peter)
	if peterRank != 3 {
		t.Errorf("expected 3, actual %d", peterRank)
		return
	}
}

// scenario C
func TestC(t *testing.T) {
	frank := &person{
		name:  "Frank",
		score: 250,
	}
	peter := &person{
		name:  "Peter",
		score: 200,
	}
	mary := &person{
		name:  "Peter",
		score: 250,
	}
	john := &person{
		name: "John",
	}
	leaderboard := &leaderboard{}
	leaderboard.add(frank)
	leaderboard.add(peter)
	leaderboard.add(mary)
	leaderboard.add(john)
	johnRank := leaderboard.getPersonRank(john)
	if johnRank != 4 {
		t.Errorf("expected 4, actual %d", johnRank)
		return
	}
}
