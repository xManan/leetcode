package solutions

func SingleNumber(nums []int) int {
    num := 0
    for _, n := range nums {
        num ^= n
    }
    return num
}
