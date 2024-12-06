package day6

type direction struct {
	onRow    int
	onColumn int
}

const (
	empty    tileType = 0
	obstacle tileType = 1
	visited  tileType = 2
)

type tileType int

var (
	up    direction = direction{onRow: -1, onColumn: 0}
	down  direction = direction{onRow: 1, onColumn: 0}
	left  direction = direction{onRow: 0, onColumn: -1}
	right direction = direction{onRow: 0, onColumn: 1}
)

type agent struct {
	rowPosition    int
	columnPosition int
	moveDirection  direction
}

func MakeMatrix(w, h int) Matrix           { return Matrix{w, h, make([]tileType, w*h)} }
func (m *Matrix) At(x, y int) tileType     { return m.data[y*m.w+x] }
func (m *Matrix) Set(x, y int, t tileType) { m.data[y*m.w+x] = t }

type environment struct {
	mapMatrix Matrix
	a         agent
}

type Matrix struct {
	w, h int
	data []tileType
}

func Part1(input []string) int {
	env := parseEnvironment(input)
	return traverseEnv(env)
}

func Part2(input []string) int {

	return 0
}

func parseEnvironment(input []string) environment {
	env := environment{}
	env.mapMatrix = MakeMatrix(len(input), len(input[0]))
	for row, line := range input {
		for column, elem := range line {
			switch elem {
			case '.':
				env.mapMatrix.Set(column, row, empty)
			case '#':
				env.mapMatrix.Set(column, row, obstacle)
			case '^', '>', 'v', '<':
				a := agent{row, column, getMoveDirection(elem)}
				env.a = a
			}
		}
	}
	return env
}

func getMoveDirection(r rune) direction {
	switch r {
	case '^':
		return up
	case '>':
		return right
	case '<':
		return left
	case 'v':
		return down
	}
	return direction{-1, -1}
}

func traverseEnv(env environment) int {
	uniqueTilesCount := 1
	for {
		if !canDoStep(env) {
			env.a.moveDirection = rotate(env.a)
		}
		hasBeenVisited, isOutside := doStep(&env)
		if isOutside {
			break
		}
		if !hasBeenVisited {
			uniqueTilesCount++
		}
	}
	return uniqueTilesCount
}

func canDoStep(env environment) bool {
	nextPosColumn := env.a.columnPosition + env.a.moveDirection.onColumn
	nextPosRow := env.a.rowPosition + env.a.moveDirection.onRow
	if nextPosColumn < 0 || nextPosColumn >= env.mapMatrix.w || nextPosRow < 0 || nextPosRow >= env.mapMatrix.h {
		return true
	}
	return env.mapMatrix.At(nextPosColumn, nextPosRow) != obstacle
}

func rotate(a agent) direction {
	currentDirection := a.moveDirection
	switch currentDirection {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return direction{-1, -1}
}

func doStep(env *environment) (alreadyVisited bool, isOutside bool) {
	a := env.a
	env.mapMatrix.Set(a.columnPosition, a.rowPosition, visited)
	a.columnPosition += a.moveDirection.onColumn
	a.rowPosition += a.moveDirection.onRow
	env.a = a
	if env.a.columnPosition < 0 || env.a.columnPosition >= env.mapMatrix.w || env.a.rowPosition < 0 || env.a.rowPosition >= env.mapMatrix.h {
		isOutside = true
		alreadyVisited = false
		return
	}
	alreadyVisited = env.mapMatrix.At(a.columnPosition, a.rowPosition) == visited
	return
}
