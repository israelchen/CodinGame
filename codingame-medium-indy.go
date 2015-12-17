package main

import "fmt"
import "os"
import "bufio"

import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Direction int

const (
	Undefined Direction = iota
	Left
	Right
	Top
	Bottom
)

type Point struct {
	X, Y int
}

type Pathway struct {
	Entry, Exit Direction
}

type Tile []Pathway

var (
	tile0  Tile = Tile{}
	tile1  Tile = Tile{{Top, Bottom}, {Right, Bottom}, {Left, Bottom}}
	tile2  Tile = Tile{{Left, Right}, {Right, Left}}
	tile3  Tile = Tile{{Top, Bottom}}
	tile4  Tile = Tile{{Top, Left}, {Right, Bottom}}
	tile5  Tile = Tile{{Top, Right}, {Left, Bottom}}
	tile6  Tile = Tile{{Left, Right}, {Right, Left}}
	tile7  Tile = Tile{{Top, Bottom}, {Right, Bottom}}
	tile8  Tile = Tile{{Left, Bottom}, {Right, Bottom}}
	tile9  Tile = Tile{{Top, Bottom}, {Left, Bottom}}
	tile10 Tile = Tile{{Top, Left}}
	tile11 Tile = Tile{{Top, Right}}
	tile12 Tile = Tile{{Right, Bottom}}
	tile13 Tile = Tile{{Left, Bottom}}

	tiles []Tile = []Tile{tile0, tile1, tile2, tile3, tile4, tile5, tile6, tile7, tile8, tile9, tile10, tile11, tile12, tile13}
)

func getNextPos(p Point, d Direction) Point {

	r := Point{p.X, p.Y}

	switch d {
	case Left:
		r.X -= 1
	case Right:
		r.X += 1
	case Top:
		r.Y -= 1
	case Bottom:
		r.Y += 1
	}

	return r
}

func tryEnterTile(entry Direction, tile Tile) (success bool, exit Direction) {

	success = false
	exit = Undefined

	for _, t := range tile {

		fmt.Fprintln(os.Stderr, "Possible: ", t)

		if t.Entry == entry {

			fmt.Fprintln(os.Stderr, "Selected: ", t)

			success = true
			exit = t.Exit
			return
		}
	}

	return success, exit
}

func StringToDirection(s string) Direction {
	switch s {
	case "TOP":
		return Top
	case "BOTTOM":
		return Bottom
	case "LEFT":
		return Left
	case "RIGHT":
		return Right
	default:
		return Undefined
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// W: number of columns.
	// H: number of rows.
	var W, H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &W, &H)

	fmt.Fprintln(os.Stderr, "W, H ", W, H)

	scene := make([][]Tile, H)

	for i := 0; i < H; i++ {
		scanner.Scan()
		LINE := scanner.Text() // represents a line in the grid and contains W integers. Each integer represents one room of a given type.

		scene[i] = make([]Tile, W)

		fmt.Fprintln(os.Stderr, "LINE", LINE)

		for j, s := range strings.Split(LINE, " ") {

			t, _ := strconv.Atoi(s)
			scene[i][j] = tiles[t]
		}
	}
	// EX: the coordinate along the X axis of the exit (not useful for this first mission, but must be read).
	var EX int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &EX)

	for {
		var XI, YI int
		var POS string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &XI, &YI, &POS)

		p := Point{XI, YI}
		t := scene[YI][XI]
		entry := StringToDirection(POS)

		if success, exit := tryEnterTile(entry, t); success {
			p = getNextPos(p, exit)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		fmt.Printf("%d %d\n", p.X, p.Y) // One line containing the X Y coordinates of the room in which you believe Indy will be on the next turn.
	}
}
