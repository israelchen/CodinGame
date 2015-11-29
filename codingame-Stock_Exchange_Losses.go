package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var err error
	var nstr string
	var n int

	reader := bufio.NewReader(os.Stdin)

	if nstr, err = reader.ReadString('\n'); err != nil {
		fmt.Fprintln(os.Stderr, "ReadString!", err)
		return
	}

	fmt.Fprintln(os.Stderr, nstr)

	if n, err = strconv.Atoi(nstr[:len(nstr)-1]); err != nil {
		fmt.Fprintln(os.Stderr, "Atoi!", err)
		return
	}

	prices := make([]int, n)

	for i := 0; i < n; i++ {

		nstr, err = reader.ReadString(' ')

		nstr = nstr[:len(nstr)-1]

		var price int

		if len(nstr) == 0 {
			continue
		}

		if price, err = strconv.Atoi(nstr); err != nil {
			fmt.Fprintln(os.Stderr, "ParseInt nstr!", err)
			continue
		}

		prices[i] = price

		fmt.Fprintf(os.Stderr, "Read: %d\n", price)
	}

	fmt.Fprintf(os.Stderr, "Read %d prices\n", n)

	maxDelta := -1
	max := prices[0]

	n = 0

	for _, price := range prices {

		delta := max - price

		if delta > maxDelta {
			maxDelta = delta
			fmt.Fprintf(os.Stderr, "Updated maxDelta: %d\n", maxDelta)
		}

		if max < price {
			max = price
			fmt.Fprintf(os.Stderr, "Updated max: %d\n", max)
		}

		n++
	}

	fmt.Fprintf(os.Stderr, "n: %d\n", n)

	if maxDelta > 0 {
		fmt.Printf("-%d\n", maxDelta) // Write answer to stdout
	} else {
		fmt.Println("0")
	}
}
