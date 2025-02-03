package solutions

func CountServers(grid [][]int) int {
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
