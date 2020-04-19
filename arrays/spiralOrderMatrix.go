/*

Getting Started with Golang, Hands On

https://www.youtube.com/watch?v=Mv_nvXXtp9E&list=PLrpeILLuqPL0vq5Al54DVDQbr8rmm8fPg

*/
func spiralOrder(matrix [][]int) []int {
    var rows int = len(matrix)    
    if rows == 0 {
        var temp []int
        return temp
    }
    var cols int = len(matrix[0])
    
    var top int = 0
    var left int = 0
    var bottom int = rows - 1
    var right int = cols - 1
    var flag int = 0
    var i, j int
    arr := make([]int, 0)
    
    for top <= bottom && left <= right {
        if flag == 0 {
            j = left
            for j <= right {
                arr = append(arr, matrix[top][j])
                j++
            }
            top++
        } else if flag == 1 {            
            i = top
            for i <= bottom {
                arr = append(arr, matrix[i][right])
                i++
            }
            right--
        } else if flag == 2 {
            j = right
            for j >= left {
                arr = append(arr, matrix[bottom][j]) 
                j--
            }
            bottom--            
        } else {
            i = bottom
            for i >= top {
                arr = append(arr, matrix[i][left])                
                i--
            }
            left++
        }
        flag = (flag + 1) % 4
    }
    return arr
    
}
