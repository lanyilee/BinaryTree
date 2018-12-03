package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

func (t *tree) deleteNode(num int) {
	var fatherTree *tree
	for {
		if t == nil {
			break
		}
		if num == t.data {
			fmt.Println("father node :" + strconv.Itoa(fatherTree.data))
			// find the data to be deleted
			if t.leftTree == nil && t.rightTree == nil {
				if fatherTree.leftTree == t {
					fatherTree.leftTree = nil //delete
				} else if fatherTree.rightTree == t {
					fatherTree.rightTree = nil //delete
				}
			} else if t.leftTree != nil && t.rightTree == nil {
				if fatherTree.leftTree == t {
					fatherTree.leftTree = t.leftTree // the deleted data contains only the left node
				} else if fatherTree.rightTree == t {
					fatherTree.rightTree = t.leftTree
				}
			} else if t.leftTree == nil && t.rightTree != nil {
				if fatherTree.leftTree == t {
					fatherTree.leftTree = t.rightTree //the deleted data contains only the right node
				} else if fatherTree.rightTree == t {
					fatherTree.rightTree = t.rightTree
				}
			} else if t.leftTree != nil && t.rightTree != nil {
				// find the minimum data in the deleted rightTree
				minTree := t.rightTree //the minimum data
				minFatherTree := t     //the father node of the minimum data
				// just follow the left path to the last node
				for {
					if minTree.leftTree != nil {
						minFatherTree = minTree
						minTree = minTree.leftTree
					} else {
						break
					}
				}
				// the leftTree of the parent of the smallest node must nil or the rightTree of the node to be deleted
				if minTree.rightTree != nil {
					minFatherTree.leftTree = minTree.rightTree
				} else {
					minFatherTree.leftTree = nil
				}

				if fatherTree.leftTree == t {
					fatherTree.leftTree = minTree
				} else if fatherTree.rightTree == t {
					fatherTree.rightTree = minTree
				}
				minTree.leftTree = t.leftTree
				minTree.rightTree = t.rightTree
			}
			break
		} else if num < t.data {
			fatherTree = t
			t = t.leftTree
		} else if num > t.data {
			fatherTree = t
			t = t.rightTree
		}
	}
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
	t.deleteNode(27)
	t.infixOrder()
}
