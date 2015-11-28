package main

import "fmt"
import "math"
import "sort"

/**
 * Save humans, destroy zombies!
 **/

// zombie to human distance
type zombieToHuman struct {
	zombieId, humanId, distance, turnsToReach int
}

// ash to human distance
type ashToHuman struct {
	humanId, distance, turnsToReach int
}

type zombie struct {
	x, y, nextx, nexty int
}

type human struct {
	x, y int
}

func calcDistance(x1, y1, x2, y2 int) int {
	// d = sqrt( sqr(x2-x1) + sqr(y2 - y1) )
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

type ZombiesToHumans []zombieToHuman

func (zth ZombiesToHumans) Len() int {
	return len(zth)
}

func (zth ZombiesToHumans) Less(i, j int) bool {
	return zth[i].turnsToReach < zth[j].turnsToReach
}

func (zth ZombiesToHumans) Swap(i, j int) {
	zth[i], zth[j] = zth[j], zth[i]
}

func main() {

	for {

		humans := make(map[int]human)
		zombies := make(map[int]zombie)

		var x, y int
		fmt.Scan(&x, &y)

		var humanCount int
		fmt.Scan(&humanCount)

		for i := 0; i < humanCount; i++ {
			var humanId, humanX, humanY int
			fmt.Scan(&humanId, &humanX, &humanY)

			humans[humanId] = human{
				humanX,
				humanY,
			}
		}
		var zombieCount int
		fmt.Scan(&zombieCount)

		for i := 0; i < zombieCount; i++ {
			var zombieId, zombieX, zombieY, zombieXNext, zombieYNext int
			fmt.Scan(&zombieId, &zombieX, &zombieY, &zombieXNext, &zombieYNext)

			zombies[zombieId] = zombie{
				zombieX,
				zombieY,
				zombieXNext,
				zombieYNext,
			}
		}

		var zombiesToHumans ZombiesToHumans = make(ZombiesToHumans, 0)
		ashToHumans := make(map[int]ashToHuman)

		for hk, hv := range humans {

			d := calcDistance(hv.x, hv.y, x, y)
			ashToHumans[hk] = ashToHuman{
				humanId:      hk,
				distance:     d,
				turnsToReach: (d - 2000) / 1000,
			}

			for zk, zv := range zombies {
				d = calcDistance(hv.x, hv.y, zv.x, zv.y)

				zombiesToHumans = append(zombiesToHumans, zombieToHuman{
					humanId:      hk,
					zombieId:     zk,
					distance:     d,
					turnsToReach: (d - 400) / 400,
				})
			}
		}

		sort.Sort(zombiesToHumans)

		h := zombiesToHumans[0]

		for _, zth := range zombiesToHumans {

			if zth.turnsToReach >= ashToHumans[zth.humanId].turnsToReach {
				h = zth
				break
			}
		}

		fmt.Printf("%d %d\n", zombies[h.zombieId].nextx, zombies[h.zombieId].nexty) // Your destination coordinates
	}
}
