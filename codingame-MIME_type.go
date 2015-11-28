package main

import "fmt"
import "os"
import "bufio"
import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&Q)
    
    mimes := make(map[string]string)
    
    for i := 0; i < N; i++ {
        // EXT: file extension
        // MT: MIME type.
        var EXT, MT string
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&EXT, &MT)
        mimes[strings.ToLower(EXT)] = MT
    }
    for i := 0; i < Q; i++ {
        
        scanner.Scan()
        FNAME := strings.ToLower(scanner.Text())
        
        idx := strings.LastIndex(FNAME, ".")
        
        if idx == -1 {
            fmt.Println("UNKNOWN")
            continue
        }
        
        ext := FNAME[idx+1:]
        
        if value, ok := mimes[ext]; ok {
            fmt.Println(value)
            continue
        }
        
        fmt.Println("UNKNOWN")
    }
}
