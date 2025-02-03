package main

import (
	"fmt"
	_ "time"

	"github.com/xManan/leetcode/solutions"
)

func main() {
    testCases := []string {
        "A man, a plan, a canal: Panama",
        "race a car",
        " ",
    }
    for _, testCase := range testCases {
	    fmt.Println(solutions.IsPalindrome(testCase))
    }
}
