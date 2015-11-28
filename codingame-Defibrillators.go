package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type defib struct {
	name        string
	address     string
	phoneNumber string
	long        float64
	lat         float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var LON string
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &LON)

	var LAT string
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &LAT)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	defibs := make(map[string]defib)
	commaReplacer := strings.NewReplacer(",", ".")

	for i := 0; i < N; i++ {
		scanner.Scan()
		fields := strings.Split(scanner.Text(), ";")

		defibLong, _ := strconv.ParseFloat(commaReplacer.Replace(fields[4]), 64)
		defibLat, _ := strconv.ParseFloat(commaReplacer.Replace(fields[5]), 64)

		newDefib := defib{
			name:        fields[1],
			address:     fields[2],
			phoneNumber: fields[3],
			long:        defibLong,
			lat:         defibLat,
		}

		defibs[fields[0]] = newDefib
	}

	minD := 100000000.0
	minName := ""

	long, _ := strconv.ParseFloat(commaReplacer.Replace(LON), 64)
	lat, _ := strconv.ParseFloat(commaReplacer.Replace(LAT), 64)

	fmt.Fprintln(os.Stderr, long)
	fmt.Fprintln(os.Stderr, lat)

	for k := range defibs {

		def := defibs[k]
		fmt.Fprintln(os.Stderr, def)

		x := (long - def.long) * math.Cos((lat+def.lat)/2.0)
		y := (lat - def.lat)
		d := 6371 * math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2))

		// fmt.Fprintln(os.Stderr, def.name, x, y, d)

		if d < minD {
			minD = d
			minName = def.name

		}
	}

	fmt.Println(minName) // Write answer to stdout
}
