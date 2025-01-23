package main

import (
	"fmt"
	_ "time"
)

func main() {
    var testCase [][]int
	testCase = [][]int{
		{1, 0},
		{0, 1},
	}
	testCase = [][]int{
		{1, 0},
		{1, 1},
	}
	testCase = [][]int{
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
    // for _,c := range testCase {
    //     fmt.Println(c)
    // }
	// fmt.Println("------------")
	fmt.Println(countServers(testCase))
}

// 1765
func highestPeak(isWater [][]int) [][]int {
	heightMap := make([][]int, len(isWater))
	var queue [][]int
	visited := make([][]int, len(isWater))
	for i, row := range isWater {
		heightMap[i] = make([]int, len(row))
		visited[i] = make([]int, len(row))
		for j, cell := range row {
			if cell == 1 {
				queue = append(queue, []int{i, j, 0})
			}
		}
	}
	for len(queue) > 0 {
        // fmt.Println(queue)
        // time.Sleep(1 * time.Second)
		src := queue[0]
		queue = queue[1:]
		if src[0]-1 >= 0 && isWater[src[0]-1][src[1]] == 0 && visited[src[0]-1][src[1]] == 0 {
            queue = append(queue, []int{src[0]-1, src[1], src[2]+1})
            visited[src[0]-1][src[1]] = 1;
        }
		if src[0]+1 < len(isWater) && isWater[src[0]+1][src[1]] == 0 && visited[src[0]+1][src[1]] == 0 {
            queue = append(queue, []int{src[0]+1, src[1], src[2]+1})
            visited[src[0]+1][src[1]] = 1
		}
		if src[1]-1 >= 0 && isWater[src[0]][src[1]-1] == 0 && visited[src[0]][src[1]-1] == 0 {
            queue = append(queue, []int{src[0], src[1]-1, src[2]+1})
            visited[src[0]][src[1]-1] = 1
        }
		if src[1]+1 < len(isWater[0]) && isWater[src[0]][src[1]+1] == 0 && visited[src[0]][src[1]+1] == 0 {
            queue = append(queue, []int{src[0], src[1]+1, src[2]+1})
            visited[src[0]][src[1]+1] = 1
		}

        heightMap[src[0]][src[1]] = src[2]
	}
	return heightMap
}

// 1267
func countServers(grid [][]int) int {
    var inRow [250][3]int
    var inCol [250][3]int

    count := 0
    var isCounted [250][250]int
    for i, row := range grid {
        for j, cell := range row {
            if cell == 1 {
                if inRow[i][0] == 1 {
                    if isCounted[i][j] != 1 {
                        count++
                        isCounted[i][j] = 1
                    }
                    if isCounted[inRow[i][1]][inRow[i][2]] != 1 {
                        count++
                        isCounted[inRow[i][1]][inRow[i][2]] = 1
                    }
                } else {
                    inRow[i] = [3]int{1, i, j}
                }
                if inCol[j][0] == 1 {
                    if isCounted[i][j] != 1 {
                        count++
                        isCounted[i][j] = 1
                    }
                    if isCounted[inCol[j][1]][inCol[j][2]] != 1 {
                        count++
                        isCounted[inCol[j][1]][inCol[j][2]] = 1
                    }
                } else {
                    inCol[j] = [3]int{1, i, j}
                }
            }
        }
    }
    return count
}
