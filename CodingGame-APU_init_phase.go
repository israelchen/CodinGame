package main

import "fmt"
import "os"
import "bufio"

//import "strings"
//import "strconv"

/**
 * Don't let the machines win. You are humanity's last hope...
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// width: the number of cells on the X axis
	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	// height: the number of cells on the Y axis
	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)

	fmt.Fprintln(os.Stderr, width)
	fmt.Fprintln(os.Stderr, height)

	var array [][]int = make([][]int, height)

	for i := 0; i < height; i++ {

		array[i] = make([]int, width)

		scanner.Scan()
		line := scanner.Text() // width characters, each either 0 or .

		for idx, r := range line {

			if r == '0' {
				array[i][idx] = 1
			} else {
				array[i][idx] = 0
			}
		}

		fmt.Fprintln(os.Stderr, line)
	}

	fmt.Fprintln(os.Stderr, array)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			if array[y][x] == 1 {

				node := fmt.Sprintf("%d %d", x, y)
				right := "-1 -1"
				bottom := "-1 -1"

				for nx := x + 1; nx < width; nx++ {
					if array[y][nx] == 1 {
						right = fmt.Sprintf("%d %d", nx, y)
						break
					}
				}

				for ny := y + 1; ny < height; ny++ {
					if array[ny][x] == 1 {
						bottom = fmt.Sprintf("%d %d", x, ny)
						break
					}
				}

				fmt.Printf("%s %s %s\n", node, right, bottom)
			}
		}
	}
}
