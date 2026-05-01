package main

import (
	"fmt"

	"github.com/yoshiyuki-140/leetcode/lru-cache/lrucache"
)

func main() {
	// test1
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]

	var lruCache lrucache.LRUCache

	scenarioVerb := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	scenarioVal := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}

	want := []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4}

	for i := 0; i < len(scenarioVerb); i++ {
		switch scenarioVerb[i] {

		case "LRUCache":
			lruCache = lrucache.NewLRUCache(scenarioVal[i][0])

		case "put":
			lruCache.Put(scenarioVal[i][0], scenarioVal[i][1])

		case "get":
			if lruCache.Get(scenarioVal[i][0]) == want[i] {
				fmt.Printf("test%d is passed\n", i)
			} else {
				fmt.Printf("test%d is failed\n", i)
			}
		}
	}
}
