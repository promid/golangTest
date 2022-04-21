package tdd

import (
	"reflect"
	"testing"
)

func TestOrderDesc(t *testing.T) {
	var inputUsers []User = []User{
		{name: "tom", maxScore: 10},
		{name: "jerry", maxScore: 11},
		{name: "jack", maxScore: 20},
		{name: "jerry-1", maxScore: 11},
	}
	var expectedUsers [][]User = [][]User{
		{
			{name: "jack", maxScore: 20},
		},
		{
			{name: "jerry", maxScore: 11},
			{name: "jerry-1", maxScore: 11},
		},
		{
			{name: "tom", maxScore: 10},
		},
	}
	topNum := 10
	var topUsers [][]User
	var err error
	topUsers, err = getTopUsersDesc(inputUsers, topNum)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(expectedUsers, topUsers) {
		t.Fatal("expected is not same as top got")
	}
}

func TestGetUserPosition(t *testing.T) {
	var users []User = []User{
		{name: "tom", maxScore: 10},
		{name: "jerry", maxScore: 11},
		{name: "jack", maxScore: 20},
		{name: "jerry-1", maxScore: 11},
	}
	var expectedPosition int = 2
	var name string = "tom"
	position := getUserPosition(users, name)
	if position != expectedPosition {
		t.Fatal("user position is not as expected")
	}
}
