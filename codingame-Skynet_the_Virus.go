package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type node struct {
	id    int
	edges []*node

	distance  int
	isGateway bool
	parent    *node
}

const InfiniteDistance int = 999999

func main() {

	g := make(map[int]*node)

	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)

	for i := 0; i < N; i++ {

		g[i] = &node{
			id:        i,
			edges:     make([]*node, 0),
			isGateway: false,
			distance:  InfiniteDistance,
			parent:    nil,
		}
	}

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)

		g[N1].edges = append(g[N1].edges, g[N2])
		g[N2].edges = append(g[N2].edges, g[N1])
	}

	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)

		g[EI].isGateway = true
	}

	for {

		for i := 0; i < len(g); i++ {
			g[i].distance = InfiniteDistance
			g[i].parent = nil
		}

		queue := make([]*node, 0)

		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)

		v := g[SI]
		v.distance = 0
		v.parent = nil

		queue = append(queue, v)

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]

			for _, adj := range u.edges {

				if adj.distance == InfiniteDistance {

					adj.distance = u.distance + 1
					adj.parent = u

					queue = append(queue, adj)
				}
			}
		}

		minDistance := InfiniteDistance
		minGateway := 0

		for i := 0; i < len(g); i++ {

			if g[i].isGateway && g[i].distance < minDistance {

				fmt.Fprintf(os.Stderr, "Gateway %d, distance: %d\n", i, g[i].distance)

				minDistance = g[i].distance
				minGateway = i
			}
		}

		fmt.Fprintf(os.Stderr, "Min gateway %d, distance: %d\n", minGateway, minDistance)

		edgeFrom := g[minGateway].id
		edgeTo := g[minGateway].parent.id

		fmt.Fprintf(os.Stderr, "Cutting edge %d <-> %d\n", edgeFrom, edgeTo)

		i := 0

		for _, e := range g[edgeFrom].edges {
			if e.id == edgeTo {
				break
			}

			i++
		}

		i = 0

		for _, e := range g[edgeTo].edges {
			if e.id == edgeFrom {
				break
			}

			i++
		}

		g[edgeTo].edges = g[edgeTo].edges[:i+copy(g[edgeTo].edges[i:], g[edgeTo].edges[i+1:])]

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		fmt.Printf("%d %d\n", edgeFrom, edgeTo)
	}
}
