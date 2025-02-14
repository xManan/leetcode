package solutions

func HighestPeak(isWater [][]int) [][]int {
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
