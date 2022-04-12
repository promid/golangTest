package tdd

import (
	"sort"
)

type User struct {
	name     string
	scores   []int
	maxScore int
}

// map[int]string
// 1 tom
// 2 jerry
// 2 jerry-1
// 4 ccc
func getTopUsersDesc(users []User, num int) ([][]User, error) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].maxScore > users[j].maxScore
	})
	var ret [][]User

	return ret, nil
}

func getUserPosition(users []User, name string) (position int) {
	//sorted, err := getTopUsersDesc(users, len(users))
	// todo handle error
	//fmt.Println(err)

	//for idx, item := range sorted {
	//	if name == item.name {
	//		position = idx
	//		return
	//	}
	//}
	return
}
