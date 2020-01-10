package main

import (
	"fmt"
	"os"
)

//文件读取
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

//定义点
type point struct {
	i, j int
}

//点相加
func (p point) add(r point) point {
	return point{
		i: p.i + r.i,
		j: p.j + r.j,
	}
}

//点在slice上的值
func (p point) at(grid [][]int) (int, bool) {
	ret, ok := 0, false
	if p.i < 0 || p.i > len(grid)-1 {
		ret, ok = 0, false
	} else if p.j < 0 || p.j > len(grid[p.i])-1 {
		ret, ok = 0, false
	} else {
		ret, ok = grid[p.i][p.j], true
	}

	return ret, ok
}

//查询点周边是否存在一个值等于val的点
func (p point) find(grid [][]int, val int) (point, bool) {
	for _, dir := range dirs {
		next := p.add(dir)
		if nextVal, ok := next.at(grid); ok && nextVal == val-1 {
			return next, true
		}
	}
	return point{0, 0}, false
}

//第一点四周的点，上左下右
var dirs = []point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

//从start开始向end探索，记录轨迹
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0] // 发现
		Q = Q[1:]

		if cur == end {
			break
		}

		//探索
		for _, dir := range dirs {
			next := cur.add(dir)
			if next == start {
				continue
			}
			curStep, _ := cur.at(steps)
			if nextStep, ok := next.at(steps); ok && nextStep == 0 {
				if curMaze, ok := next.at(maze); ok && curMaze == 0 {
					Q = append(Q, next)
					steps[next.i][next.j] = curStep + 1
				}
			}
		}
	}
	return steps

}

//通过探索轨迹反追踪路径
func trace(steps [][]int, p point, traces []point) {
	if val, ok := p.at(steps); ok && val > 0 {
		for i := 0; i < val; i++ {
			if pre, ok := p.find(steps, val); ok {
				val -= 1
				traces[val] = pre
				if val > 1 {
					trace(steps, pre, traces)
				}
			}
		}
	}
}

func main() {
	maze := readMaze("maze/maze.in")
	start := point{0, 0}
	end := point{5, 4}
	steps := walk(maze, start, end)

	step, _ := end.at(steps)
	traces := make([]point, step+1)
	traces[step] = end
	trace(steps, end, traces)
	traces[0] = start
	for k, p := range traces {
		fmt.Printf("%2d : %v \n", k, p)
	}
}
