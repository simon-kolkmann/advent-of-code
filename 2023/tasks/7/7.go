package task_7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct {
	utils.BaseSolver
}

type Card struct {
	symbol   string
	strength int
}

func cardFromSymbol(symbol string) Card {
	var strength int

	switch symbol {
	case "2":
		strength = 1
	case "3":
		strength = 2
	case "4":
		strength = 3
	case "5":
		strength = 4
	case "6":
		strength = 5
	case "7":
		strength = 6
	case "8":
		strength = 7
	case "9":
		strength = 8
	case "T":
		strength = 9
	case "J":
		strength = 10
	case "Q":
		strength = 11
	case "K":
		strength = 12
	case "A":
		strength = 13
	}

	return Card{
		symbol:   symbol,
		strength: strength,
	}
}

type Hand struct {
	cards []Card
	bid   int
}

func NewHand(line string) Hand {
	split := strings.Split(line, " ")
	cards := split[0]
	bid := split[1]

	hand := Hand{}

	for _, symbol := range strings.Split(cards, "") {
		hand.cards = append(hand.cards, cardFromSymbol(symbol))
	}

	sort.Slice(hand.cards, func(i, j int) bool {
		return hand.cards[i].strength > hand.cards[j].strength
	})

	hand.bid, _ = strconv.Atoi(bid)

	return hand
}

func (hand Hand) getType() int {
	stats := make(map[string]int)

	for _, card := range hand.cards {
		stats[card.symbol]++
	}

	// TODO: check stats and return type
	return 0
}

func NewSolver() Solver {
	return Solver{
		utils.NewBaseSolver(7),
	}
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	hands := make([]Hand, 0)

	for _, line := range lines {
		hands = append(hands, NewHand(line))
	}

	// TODO: solve by type & card strength
	fmt.Println(hands)

	return ""
}

func (solver Solver) SolveSecond() string {
	// lines, _ := solver.GetPuzzleInput()
	return ""
}
