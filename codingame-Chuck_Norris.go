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
 
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    MESSAGE := scanner.Text()
    
    output := ""
    binary := ""

    for _, r := range MESSAGE {
        
        bin := strconv.FormatInt(int64(r), 2)
        
        if len(bin) < 7 {
            bin = strings.Repeat("0", 7 - len(bin)) + bin
        }
        
        binary += bin
    }

    fmt.Fprintln(os.Stderr, binary)

    last := 'x'
    count := 0
    
    for _, bit := range binary {
        
        if last != bit {
            
            if count > 0 {
                output += " "
            }
            
            if bit == '1' {
                output += "0 "
            } else {
                output += "00 "
            }

            last = bit
            count = 0
        }
        
        output += "0"

        count++
    }

    fmt.Println(output)// Write answer to stdout
}
