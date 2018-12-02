package main

import (
	"fmt"
	"math/rand"
)

func main() {

	t := &tree{15, nil, nil}
	for i := 0; i < 30; i++ {
		rangeNum := rand.Int31n(30)
		fmt.Println(rangeNum)
		t.addNode(int(rangeNum))
	}
	t.infixOrder()
}
