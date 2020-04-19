'''
Getting Started with Golang, HandsOn

https://www.youtube.com/watch?v=Mv_nvXXtp9E&list=PLrpeILLuqPL0vq5Al54DVDQbr8rmm8fPg

'''
def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        rows = len(matrix)
        if rows:
            cols = len(matrix[0])  
        else:
            return matrix            
        
        left = 0
        top = 0
        right = cols - 1
        bottom = rows - 1
        flag = 0
        arr = []
        
        while left <= right and top <= bottom:            
            if flag == 0:
                # top most row
                j = left 
                while j <= right:
                    arr.append(matrix[top][j])    
                    j = j + 1
                top = top + 1
            elif flag == 1:
                # right most col
                i = top
                while i <= bottom:
                    arr.append(matrix[i][right])
                    i = i + 1
                right = right - 1
            elif flag == 2:
                # bottom most row
                j = right
                while j >= left:
                    arr.append(matrix[bottom][j])
                    j = j - 1
                bottom = bottom - 1
            else:
                # left most col
                i = bottom 
                while i >= top:
                    arr.append(matrix[i][left])
                    i = i - 1
                left = left + 1
            flag = (flag + 1) % 4
        return arr
