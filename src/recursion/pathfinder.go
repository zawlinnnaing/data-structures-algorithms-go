package recursion

import "fmt"

// Path finder recursion
/*
Base cases:
- At the end
- Hit the wall
- Already visited
- Out of the map

Recursive cases:
- Goes all four directions
*/

type Path struct {
	points []Point
}

func (path *Path) Pop() (v Point) {
	pathLength := len(path.points)
	if len(path.points) == 0 {
		return
	}
	newPoints, lastItem := path.points[:pathLength-2], path.points[pathLength-1]
	path.points = newPoints
	return lastItem
}

func (path *Path) Push(point Point) {
	path.points = append(path.points, point)
}

type Point struct {
	x int
	y int
}

var dir [4][2]int = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func walk(maze [][]string, wall string, currentPoint Point, end Point, seen [][]bool, path *Path) bool {
	// Base case: off the map
	if currentPoint.x >= len(maze[0]) || currentPoint.x < 0 || currentPoint.y < 0 || currentPoint.y >= len(maze) {
		return false
	}
	// Base case: hit the wall
	if maze[currentPoint.y][currentPoint.x] == wall {
		return false
	}

	// Arrive at end
	if currentPoint.x == end.x && currentPoint.y == end.y {
		path.Push(currentPoint)
		return true
	}

	// Base case: already visited
	if seen[currentPoint.y][currentPoint.x] {
		return false
	}

	seen[currentPoint.y][currentPoint.x] = true
	path.Push(currentPoint)

	for i := 0; i < len(dir); i++ {
		if walk(maze, wall, Point{
			x: currentPoint.x + dir[i][1],
			y: currentPoint.y + dir[i][0],
		}, end, seen, path) {
			return true
		}
	}

	path.Pop()

	return false
}

func PathFinder(maze [][]string, wall string, start Point, end Point) Path {
	path := Path{points: []Point{}}
	seen := make([][]bool, 0)

	for i := 0; i < len(maze[0]); i++ {
		innerSeen := make([]bool, len(maze[0]))
		seen = append(seen, innerSeen)
	}
	walk(maze, wall, start, end, seen, &path)

	return path
}

func PathFinderExample() {
	wall := "#"
	maze := [][]string{
		{"#", "#", "e", "#"},
		{"#", "p", "p", "#"},
		{"#", "p", "p", "#"},
		{"#", "s", "#", "#"},
	}
	start := Point{
		x: 1,
		y: 3,
	}
	end := Point{
		x: 2,
		y: 0,
	}
	path := PathFinder(maze, wall, start, end)

	fmt.Println("Found path:", path.points)
}
