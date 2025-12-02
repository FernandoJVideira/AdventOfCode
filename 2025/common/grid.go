package common

// Points
type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Neighbors4() (res []Point) {
	for _, dir := range DirectionVectors {
		res = append(res, p.Add(dir))
	}
	return res
}

func (p Point) Neighbors8() (res []Point) {
	for _, dir := range AllDirectionVectors {
		res = append(res, p.Add(dir))
	}
	return res
}

func (p Point) ManhattanDistance(otherPoint Point) int {
	return Abs(p.X-otherPoint.X) + Abs(p.Y-otherPoint.Y)
}

// Directions
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var DirectionVectors = []Point{
	{X: 0, Y: -1}, // North
	{X: 1, Y: 0},  // East
	{X: 0, Y: 1},  // South
	{X: -1, Y: 0}, // West
}

var AllDirectionVectors = []Point{
	{X: 0, Y: -1},  // N
	{X: 1, Y: -1},  // NE
	{X: 1, Y: 0},   // E
	{X: 1, Y: 1},   // SE
	{X: 0, Y: 1},   // S
	{X: -1, Y: 1},  // SW
	{X: -1, Y: 0},  // W
	{X: -1, Y: -1}, // NW
}

func (d Direction) TurnRight() Direction {
	return Direction((int(d) + 1) % 4)
}

func (d Direction) TurnLeft() Direction {
	return Direction((int(d) + 3) % 4)
}

// Grid
type Grid [][]rune

func NewGrid(lines []string) Grid {
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	if g.Height() == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid) InBounds(p Point) bool {
	return p.Y >= 0 && p.Y < g.Height() && p.X >= 0 && p.X < g.Width()
}

func (g Grid) Get(p Point) rune {
	// 1. Check if the point is in bounds
	if g.InBounds(p) {
		return g[p.Y][p.X]
	}
	return 0
}

func (g Grid) Set(p Point, value rune) {
	if g.InBounds(p) {
		g[p.Y][p.X] = value
	}
}

func (g Grid) Find(target rune) (Point, bool) {
	for y, row := range g {
		for x, val := range row {
			if val == target {
				return Point{X: x, Y: y}, true
			}
		}
	}
	return Point{}, false
}

func (g Grid) FindAll(target rune) (res []Point) {
	for y, row := range g {
		for x, val := range row {
			if val == target {
				res = append(res, Point{X: x, Y: y})
			}
		}
	}
	return res
}

func (g Grid) BreadthFirstSearch(start, goal Point, wall rune) ([]Point, bool) {
	if !g.InBounds(start) || !g.InBounds(goal) {
		return nil, false
	}

	cameFrom := make(map[Point]Point)
	queue := []Point{start}

	cameFrom[start] = start

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == goal {
			// Reconstruct path
			path := []Point{}
			for p := current; p != start; p = cameFrom[p] {
				path = append([]Point{p}, path...)
			}
			path = append([]Point{start}, path...)
			return path, true
		}

		for _, neighbor := range current.Neighbors4() {
			if g.InBounds(neighbor) && g.Get(neighbor) != wall {
				if _, visited := cameFrom[neighbor]; !visited {
					cameFrom[neighbor] = current
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return nil, false
}
