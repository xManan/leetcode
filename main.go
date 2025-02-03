package main

import (
	"fmt"
	_ "time"

	"github.com/xManan/leetcode/solutions"
)

func main() {
    testCases := [][]int {
        {1,4,3,3,2},
        {3,3,3,3},
        {3,2,1},
    }
    for _, testCase := range testCases {
	    fmt.Println(solutions.LongestMonotonicSubarray(testCase))
    }
}
