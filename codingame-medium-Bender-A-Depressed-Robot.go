package main

import "fmt"
import "os"
import "bufio"
import "strings"

//import "strconv"

type Point struct {
	X, Y int
}

type Direction int

const (
	DirectionNorth Direction = iota
	DirectionSouth
	DirectionEast
	DirectionWest
)

var directions map[Direction]Point = map[Direction]Point{
	DirectionNorth: Point{0, -1},
	DirectionSouth: Point{0, 1},
	DirectionEast:  Point{1, 0},
	DirectionWest:  Point{-1, 0},
}

var directionNames map[Direction]string = map[Direction]string{
	DirectionNorth: "NORTH",
	DirectionSouth: "SOUTH",
	DirectionEast:  "EAST",
	DirectionWest:  "WEST",
}

type Path struct {
	pos       Point
	dir       Direction
	isBreaker bool
}

var directionFinding []Direction = []Direction{DirectionSouth, DirectionEast, DirectionNorth, DirectionWest}
var directionFindingInv []Direction = []Direction{DirectionWest, DirectionNorth, DirectionEast, DirectionSouth}

func (p Point) Move(dir Direction) Point {

	p2 := directions[dir]
	p3 := Point{p.X + p2.X, p.Y + p2.Y}

	//fmt.Fprintln(os.Stderr, "p1: ", p, ", p2: ", p2, "p3: ", p3)

	return p3
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var L, C int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L, &C)

	board := make([][]byte, L)
	teleports := make([]Point, 2)

	var startPos, endPos Point
	teleport := 0

	for i := 0; i < L; i++ {
		scanner.Scan()
		row := scanner.Text()

		if j := strings.Index(row, "@"); j != -1 {
			startPos = Point{j, i}
		}

		if j := strings.Index(row, "$"); j != -1 {
			endPos = Point{j, i}
		}

		if j := strings.Index(row, "T"); j != -1 {
			teleports[teleport] = Point{j, i}
			teleport++
		}

		board[i] = make([]byte, C)

		for j := 0; j < len(row); j++ {
			board[i][j] = row[j]
		}

		fmt.Fprintln(os.Stderr, row)
	}

	history := []Path{}
	pos := startPos

	dirFindingRule := directionFinding
	dirFindingIndex := 0

	isBlocked, isInverted, isBreaker, isDirFinding, isLoop := false, false, false, true, false
	var dir, newDir Direction

	step := 0

	for {

		if pos == endPos {
			break
		}

		if isDirFinding {
			dir = dirFindingRule[dirFindingIndex]
			dirFindingIndex++
		}

		next := pos.Move(dir)
		ch := board[next.Y][next.X]

		newDir = dir

		isBlocked = false

		switch ch {
		case 'B':
			isBreaker = !isBreaker

		case '#':
			isBlocked = true

		case 'X':
			if isBreaker {
				board[next.Y][next.X] = ' '
			} else {
				isBlocked = true
			}

		case 'I':
			isInverted = !isInverted

			if isInverted {
				dirFindingRule = directionFindingInv
			} else {
				dirFindingRule = directionFinding
			}

		case 'S':
			newDir = DirectionSouth

		case 'N':
			newDir = DirectionNorth

		case 'E':
			newDir = DirectionEast

		case 'W':
			newDir = DirectionWest

		case 'T':

			if next == teleports[0] {
				next = teleports[1]
			} else {
				next = teleports[0]
			}

			fmt.Fprintln(os.Stderr, "teleported to: ", next)
		}

		if isBlocked {
			isDirFinding = true
			continue
		}

		pathRecord := Path{pos, dir, isBreaker}

		if step > L*C {
			isLoop = true
			break
		}

		history = append(history, pathRecord)

		fmt.Fprintln(os.Stderr, "step: ", step, "dir: ", directionNames[dir], "newDir: ", directionNames[newDir], "pos: ", pos, "next: ", next, "isBreaker: ", isBreaker, "isInverted: ", isInverted, "isLoop: ", isLoop)

		dir = newDir
		step++
		pos = next
		isDirFinding = false
		dirFindingIndex = 0
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	if isLoop {
		fmt.Println("LOOP")
	} else {
		for _, path := range history {
			fmt.Println(directionNames[path.dir]) // Write answer to stdout
		}
	}
}
