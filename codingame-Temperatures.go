package main

import "fmt"
import "os"
import "bufio"
import "math"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)
    
    scanner.Scan()
    
    temps := strings.Split(scanner.Text(), " ") // the n temperatures expressed as integers ranging from -273 to 5526
    
    fmt.Fprintln(os.Stderr, temps)

    var closest int = 10000
    
    for _, t := range temps {
        
        temp, _ := strconv.Atoi(t)

        fmt.Fprintln(os.Stderr, "x", temp, closest)

        if math.Abs(float64(temp)) < math.Abs(float64(closest)) {
            fmt.Fprintln(os.Stderr, "xx", temp, closest)
            closest = temp
        } else if math.Abs(float64(temp)) == math.Abs(float64(closest)) {

            fmt.Fprintln(os.Stderr, "xxx", temp, closest)
            
            if (temp > 0 && closest < 0) {
                fmt.Fprintln(os.Stderr, "xxxx", temp, closest)
                closest = temp
            }
        }
    }

    fmt.Println(strconv.Itoa(closest))// Write answer to stdout
}
