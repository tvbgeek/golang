/*

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
 

Example 1:


Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Example 2:

Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 

*/


func isValidSudoku(board [][]byte) bool {
    rows := len(board)
    cols := len(board[0])
    dc := make(map[byte]int, 9)
    
    if rows != 9 || cols != 9 {
        return false
    }
    
    // check elements in rows
    for i := 0; i < rows; i++ {
        dc = make(map[byte]int, 9)
        for j := 0; j < cols; j++ {
            if board[i][j] != '.' {
                _, prs := dc[board[i][j]]
                if prs {
                    fmt.Printf("i %d j %d dc[board[i][j]] %d", i, j, dc[board[i][j]])
                   return false
                } else {
                    dc[board[i][j]] = 1
             }
            }
        }
    }
    
    // check elements in cols
    for j := 0; j < cols; j++ {
        dc = make(map[byte]int, 9)
        for i := 0; i < rows; i++ {
            if board[i][j] != '.' {
                _, prs := dc[board[i][j]]
                if prs {
                    fmt.Printf("i %d j %d dc[board[i][j]] %d", i, j, dc[board[i][j]])
                    return false
                } else {
                    dc[board[i][j]] = 1
                }
            }
        }
    }
    
    // check for 3x3 submetrices
    for i := 0; i < rows; i=i+3 {
        for j:= 0; j < cols; j=j+3 {
            dc = make(map[byte]int, 9)            
            for k := i; k < i+3; k++ {
                for l := j; l < j+3; l++ {
                    if board[k][l] != '.' {
                        _, prs := dc[board[k][l]]
                        if prs {
                            fmt.Printf("i %d j %d dc[board[i][j]] %d", i, j, dc[board[k][l]])
                            return false
                        } else {
                            dc[board[k][l]] = 1
                        }
                    }                    
                }
            }
        }
    }
    
    return true
}
