package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	data      int
	leftTree  *tree
	rightTree *tree
}

func (t *tree) addNode(num int) {
	if num == t.data {
		return
	} else if num < t.data {
		if t.leftTree == nil {
			t.leftTree = &tree{num, nil, nil}
			return
		} else {
			t = t.leftTree
		}
	} else if num > t.data {
		if t.rightTree == nil {
			t.rightTree = &tree{num, nil, nil}
			return
		} else {
			t = t.rightTree
		}
	}
	//recursion
	t.addNode(num)
}

func (t *tree) searchNode(num int) {
	for {
		if t == nil {
			break
		}
		if num == t.data {
			fmt.Println("find it")
			return
		} else if num < t.data {
			t = t.leftTree
		} else if num > t.data {
			t = t.rightTree
		}
	}
	fmt.Println("can not find it")
}

func (t *tree) infixOrder() {
	if t == nil {
		return
	}
	left := t.leftTree
	left.infixOrder()
	fmt.Println(t.data)
	right := t.rightTree
	right.infixOrder()
}

func (t *tree) addNode2(num int) {
	addTree := &tree{num, nil, nil}
	for {
		if num == t.data {
			break
		} else if num < t.data {
			if t.leftTree == nil {
				t.leftTree = addTree
				break
			} else {
				t = t.leftTree
			}
		} else if num > t.data {
			if t.rightTree == nil {
				t.rightTree = addTree
				break
			} else {
				t = t.rightTree
			}
		}
	}
}

func main() {

	t := &tree{15, nil, nil}
	for i := 0; i < 30; i++ {
		fmt.Println("rangeNum:")
		rangeNum := rand.Int31n(30)
		fmt.Println(rangeNum)
		t.addNode(int(rangeNum))
	}
	t.infixOrder()
	t.searchNode(15)
}
