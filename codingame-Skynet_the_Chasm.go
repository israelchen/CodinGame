package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // road: the length of the road before the gap.
    var road int
    fmt.Scan(&road)
    
    // gap: the length of the gap.
    var gap int
    fmt.Scan(&gap)
    
    // platform: the length of the landing platform.
    var platform int
    fmt.Scan(&platform)

    for {
        // speed: the motorbike's speed.
        var speed int
        fmt.Scan(&speed)
        
        // coordX: the position on the road of the motorbike.
        var coordX int
        fmt.Scan(&coordX)
        
        
        command := "WAIT"
        
        if coordX < road && speed < gap + 1{
            command = "SPEED"
        } else if coordX < road && speed > gap + 1 {
            command = "SLOW"
        } else if coordX == road - 1 {
            command = "JUMP"
        } else if coordX >= road + gap  {
            command = "SLOW"
        }

        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        fmt.Println(command) // A single line containing one of 4 keywords: SPEED, SLOW, JUMP, WAIT.
    }
}
