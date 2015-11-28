package main

import "fmt"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // W: width of the building.
    // H: height of the building.
    var W, H int
    fmt.Scan(&W, &H)
    
    // N: maximum number of turns before game over.
    var N int
    fmt.Scan(&N)
    
    var x, y int
    fmt.Scan(&x, &y)
    
    var x0, y0, x1, y1, dx, dy int
    
    fmt.Fprintf(os.Stderr, "building: (%d,%d)\n", W, H)
    fmt.Fprintf(os.Stderr, "N: %d\n", N)
    fmt.Fprintf(os.Stderr, "batman: (%d,%d)\n", x, y)
    
    x0 = 0
    y0 = 0
    x1 = W-1
    y1 = H-1
    
    for {
        // BOMB_DIR: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        var BOMB_DIR string
        fmt.Scan(&BOMB_DIR)
        
        switch BOMB_DIR {
            case "U":
                // adjust box
                y1 = y
            
                dx = 0
                dy = int(math.Floor(float64((y0 - y)) / 2.0))
            case "UL":
                // adjust box
                y1 = y
                x1 = x
            
                dx = int(math.Floor(float64((x0 - x)) / 2.0))
                dy = int(math.Floor(float64((y0 - y)) / 2.0))
            case "UR":
                // adjust box
                y1 = y
                x0 = x
            
                dx = int(math.Ceil(float64((x1 - x)) / 2.0))
                dy = int(math.Floor(float64((y0 - y)) / 2.0))
            case "R":
                // adjust box
                x0 = x
                
                dx = int(math.Ceil(float64((x1 - x)) / 2.0))
                dy = 0
            case "D":
                // adjust box
                y0 = y
                
                dx = 0
                dy = int(math.Ceil(float64((y1 - y)) / 2.0))
            case "DL":
                // adjust box
                x1 = x
                y0 = y
                
                dx = int(math.Floor(float64((x0 - x)) / 2.0))
                dy = int(math.Ceil(float64((y1 - y)) / 2.0))
            case "DR":
                // adjust box
                x0 = x
                y0 = y
            
                dx = int(math.Ceil(float64((x1 - x)) / 2.0))
                dy = int(math.Ceil(float64((y1 - y)) / 2.0))
            case "L":
                // adjust box
                x1 = x
            
                dx = int(math.Floor(float64((x0 - x)) / 2.0))
                dy = 0
        }
        
        x += dx
        y += dy
        
        fmt.Fprintf(os.Stderr, "dir: %s, box: (%d,%d) (%d,%d)\n", BOMB_DIR, x,y,x1,y1)
        
        fmt.Fprintf(os.Stderr, "delta: (%d,%d), next: (%d,%d)\n", dx, dy, x, y)
        fmt.Printf("%d %d\n", x, y)
    }
}
