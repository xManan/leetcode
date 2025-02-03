package solutions

func LongestMonotonicSubarray(nums []int) int {
    incCount := 1
    maxIncCount := incCount
    decCount := 1
    maxDecCount := decCount
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] < nums[i+1] {
            incCount++
        } else {
            if incCount > maxIncCount {
                maxIncCount = incCount
            }
            incCount = 1
        }
        if nums[i] > nums[i+1] {
            decCount++
        } else {
            if decCount > maxDecCount {
                maxDecCount = decCount
            }
            decCount = 1
        }
    }
    if incCount > maxIncCount {
        maxIncCount = incCount
    }
    if decCount > maxDecCount {
        maxDecCount = decCount
    }
    return max(maxIncCount, maxDecCount)
}
