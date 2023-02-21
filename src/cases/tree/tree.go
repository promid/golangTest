package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type executor interface {
	execute()
}

type Tree struct {
	sync.Mutex
	sentinel *Leaf
	running  map[int]bool
}

func NewTree(sentinel *Leaf, running map[int]bool) *Tree {
	m := running
	if m == nil {
		m = make(map[int]bool)
	}
	return &Tree{sentinel: sentinel, running: m}
}

func (t *Tree) SetRunning(id int, running bool) (existing bool) {
	t.Lock()
	defer func() {
		t.running[id] = running
		t.Unlock()
	}()
	_, existing = t.running[id]
	return
}

func (t *Tree) isLeafRunning(id int) (running bool, existing bool) {
	t.Lock()
	defer t.Unlock()
	running, existing = t.running[id]
	return
}

func (t *Tree) isLeafDone(id int) bool {
	t.Lock()
	defer t.Unlock()
	running, existing := t.running[id]
	return existing && !running
}

type Leaf struct {
	// executor executor
	id       int
	value    string
	children []*Leaf
	parents  []*Leaf
}

func (l *Leaf) execute() {
	randInt := rand.Intn(10)
	fmt.Printf("leaf id %d, will run %d sec\n", l.id, randInt)
	time.Sleep(time.Duration(randInt) * time.Second)
	fmt.Printf("leaf id %d is done\n", l.id)
}

func (t *Tree) execute() {
	for _, child := range t.sentinel.children {
		child := child
		go func() {
			wait := true
			for wait {
				var done int
				var parentsIds []int
				for _, p := range child.parents {
					running, existing := t.isLeafRunning(p.id)
					if existing && !running {
						done++
					} else {
						parentsIds = append(parentsIds, p.id)
					}
				}
				wait = len(child.parents) != done
				if wait {
					time.Sleep(time.Second)
					fmt.Printf("leaf %d is waiting for parents %v\n", child.id, parentsIds)
				}
			}
			if runningBeforeSetting := t.SetRunning(child.id, true); runningBeforeSetting {
				fmt.Printf("leaf %d already running, skip\n", child.id)
				return
			}
			child.execute()
			t.SetRunning(child.id, false)
			NewTree(child, t.running).execute()
		}()
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	sentinel := &Leaf{
		id:    0,
		value: "sentinel",
	}
	leaf1 := &Leaf{
		id:    1,
		value: "1",
	}
	leaf2 := &Leaf{
		id:    2,
		value: "2",
	}
	leaf3 := &Leaf{
		id:      3,
		value:   "3",
		parents: []*Leaf{leaf1},
	}
	leaf1.children = []*Leaf{leaf3}
	leaf4 := &Leaf{
		id:      4,
		value:   "4",
		parents: []*Leaf{leaf2},
	}
	leaf2.children = []*Leaf{leaf4}
	leaf5 := &Leaf{
		id:      5,
		value:   "5",
		parents: []*Leaf{leaf3, leaf4},
	}
	leaf3.children = []*Leaf{leaf5}
	leaf4.children = []*Leaf{leaf5}
	sentinel.children = []*Leaf{leaf1, leaf2}
	tree := NewTree(sentinel, make(map[int]bool))
	tree.execute()
	time.Sleep(100 * time.Second)
}
