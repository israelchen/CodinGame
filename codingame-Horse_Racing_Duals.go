package main

import "fmt"
import "os"
import "sort"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
    fmt.Scan(&N)
    
    powers := make([]int, 0)
    
    for i := 0; i < N; i++ {
        var Pi int
        fmt.Scan(&Pi)
        
        if Pi > 0 {
            powers = append(powers, Pi)
            fmt.Fprintln(os.Stderr, Pi)
        }
    }

    fmt.Fprintln(os.Stderr, powers)
    
    sort.Ints(powers)
    
    minD := 99999999999
    
    for i := 1; i < len(powers); i++ {
        
        delta := int(math.Abs(float64(powers[i] - powers[i - 1])))
        
        if delta < minD {
            minD = delta
        }
    }
    
    fmt.Fprintln(os.Stderr, powers)
    
    fmt.Println(minD)// Write answer to stdout
}
