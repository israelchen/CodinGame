package main

import "fmt"
import "os"

type battleResult int

const (
	battle_p1_wins battleResult = iota
	battle_p2_wins
	battle_draw
	battle_p1_out_of_cards
	battle_p2_out_of_cards
)

var battleResultStrings map[battleResult]string = map[battleResult]string{
	battle_p1_wins:         "battle_p1_wins",
	battle_p2_wins:         "battle_p2_wins",
	battle_draw:            "battle_draw",
	battle_p1_out_of_cards: "battle_p1_out_of_cards",
	battle_p2_out_of_cards: "battle_p2_out_of_cards",
}

type drawResult int

const (
	draw_p1_wins drawResult = iota
	draw_p2_wins
	draw_tie
	draw_p1_out_of_cards
	draw_p2_out_of_cards
)

var (
	cardsp1, cardsp2 []string
	i, battlesFought int
	power            map[string]int = map[string]int{"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "10": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
)

func drawCards() drawResult {

	if i >= len(cardsp1) {
		return draw_p1_out_of_cards
	}

	if i >= len(cardsp2) {
		return draw_p2_out_of_cards
	}

	cardp1 := cardsp1[i]
	cardp2 := cardsp2[i]
    
	fmt.Fprintf(os.Stderr, "p1 draws: %s (len: %d)\n", cardp1, len(cardsp1))
	fmt.Fprintf(os.Stderr, "p2 draws: %s (len: %d)\n", cardp2, len(cardsp2))
    
	powerp1 := power[cardp1[0:len(cardsp1[i])-1]]
	powerp2 := power[cardp2[0:len(cardsp2[i])-1]]

    i += 1

	if powerp1 > powerp2 {
		return draw_p1_wins
	} else if powerp1 < powerp2 {
		return draw_p2_wins
	} else {
		return draw_tie
	}
}

func doWar() battleResult {

	fmt.Fprintf(os.Stderr, "War!!\n")

	for count := 0; count < 3; count++ {

		d := drawCards()

		if d == draw_p1_out_of_cards || d == draw_p2_out_of_cards {
			return battle_draw
		}
	}

	return doBattle(true)
}

func doBattle(inWar bool) battleResult {

	fmt.Fprintf(os.Stderr, "p1: %q\n", cardsp1)
	fmt.Fprintf(os.Stderr, "p2: %q\n", cardsp2)

	d := drawCards()

	switch d {
	case draw_p1_wins:

		for _, c := range cardsp1[:i] {
			cardsp1 = append(cardsp1, c)
		}

		for _, c := range cardsp2[:i] {
			cardsp1 = append(cardsp1, c)
		}

		cardsp1 = cardsp1[i:]
		cardsp2 = cardsp2[i:]

		i = 0
		battlesFought++

		return battle_p1_wins

	case draw_p2_wins:

		for _, c := range cardsp1[:i] {
			cardsp2 = append(cardsp2, c)
		}

		for _, c := range cardsp2[:i] {
			cardsp2 = append(cardsp2, c)
		}

		cardsp1 = cardsp1[i:]
		cardsp2 = cardsp2[i:]

		i = 0
		battlesFought++

		return battle_p2_wins

	case draw_tie:
		return doWar()

	case draw_p1_out_of_cards:
		if inWar {
			return battle_draw
		} else {
			return battle_p1_out_of_cards
		}

	case draw_p2_out_of_cards:
		if inWar {
			return battle_draw
		} else {
			return battle_p2_out_of_cards
		}
	}

	return battle_draw
}

func battle() battleResult {
	return doBattle(false)
}

func main() {

	var n int
	fmt.Scan(&n)

	cardsp1 = make([]string, n)

	for i := 0; i < n; i++ {

		var cardp1 string
		fmt.Scan(&cardp1)

		cardsp1[i] = cardp1
	}

	var m int
	fmt.Scan(&m)

	cardsp2 = make([]string, m)

	for i := 0; i < m; i++ {

		var cardp2 string
		fmt.Scan(&cardp2)

		cardsp2[i] = cardp2
	}

	winner := "PAT"

	battlesFought = 0
	i = 0

	for {

		result := battle()
		fmt.Fprintf(os.Stderr, "battle %d result: %s\n", battlesFought, battleResultStrings[result])

		if result == battle_p2_out_of_cards {
			winner = "1"
			break
		}

		if result == battle_p1_out_of_cards {
			winner = "2"
			break
		}

		if result == battle_draw {
			winner = "PAT"
			break
		}
	}

    if winner == "PAT" {
    	fmt.Println("PAT")
    } else {
    	fmt.Printf("%s %d\n", winner, battlesFought)
    }
}
