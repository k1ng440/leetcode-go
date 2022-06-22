package floodfill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	visited := make([][]bool, len(image))
	for i := range visited {
		visited[i] = make([]bool, len(image[i]))
	}

	value := image[sr][sc]
	dfs(visited, image, value, sr, sc, color)
	return image
}

func dfs(visited [][]bool, image [][]int, val, sr, sc, color int) {
	if sr < 0 || len(image)-1 < sr {
		return
	}

	if sc < 0 || len(image[sr])-1 < sc {
		return
	}

	if visited[sr][sc] {
		return
	}

	if image[sr][sc] != val {
		return
	}

	visited[sr][sc] = true
	image[sr][sc] = color
	dfs(visited, image, val, sr-1, sc, color) // up
	dfs(visited, image, val, sr+1, sc, color) // down
	dfs(visited, image, val, sr, sc-1, color) // left
	dfs(visited, image, val, sr, sc+1, color) // right
}
