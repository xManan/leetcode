package main

import (
	"fmt"
	_ "time"

	"github.com/xManan/leetcode/solutions"
)

func main() {
    var testCase []int
	testCase = []int{2,2,1}
	testCase = []int{4,1,2,1,2}
	testCase = []int{1}
	fmt.Println(solutions.SingleNumber(testCase))
}
