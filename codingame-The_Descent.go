package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    for {
        var spaceX, spaceY int
        fmt.Scan(&spaceX, &spaceY)
        
        max := 0
        maxX := 9
        
        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
            var mountainH int
            fmt.Scan(&mountainH)
            
            if mountainH > max {
                max = mountainH
                maxX = i
            }
        }
        
        if spaceX == maxX {
            fmt.Println("FIRE")
        } else {
            fmt.Println("HOLD")
        }
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
         // either:  FIRE (ship is firing its phase cannons) or HOLD (ship is not firing).
    }
}
