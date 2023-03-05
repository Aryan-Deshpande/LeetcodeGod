package graphs

func numIslands() int {
    var numis int = 0
	var grid = [][] int{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0},
	}
    for i := 0; i< len(grid); i++{
        for j:= 0; j< len(grid[0]); j++{
            if(grid[i][j] == 0 || grid[i][j] == 2){
                continue
            }else{
                checkIslands(0, 0, grid)
                numis=numis+1
            }
        }
    }

    return numis
}

func checkIslands(x int, y int, grid[][] int) {
    if(x < 0 || y < 0 || len(grid) >= x || len(grid[0]) >= y || grid[x][y] == 0 || grid[x][y] == 2){
        return;
    }else{
        grid[x][y] = 2;
    } 
    
    checkIslands(x+1, y, grid)
    checkIslands(x, y+1, grid)
    checkIslands(x-1, y, grid)
    checkIslands(x, y-1, grid)

    return;
}