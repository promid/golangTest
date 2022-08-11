package learboard

import (
	"sort"
)

type person struct {
	name  string
	score int
}

type leaderboard struct {
	persons map[string]int
}

func (l *leaderboard) add(p *person) {
	if l.persons == nil {
		l.persons = make(map[string]int)
	}
	l.persons[p.name] = p.score
}

func (l *leaderboard) getPersonRank(p *person) int {

	return 1
}

func sortPersonsDesc(persons []person) {
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].score > persons[j].score
	})
}
