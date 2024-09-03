package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	monthValues := GetMonthValues("halfright")
	for _, mv := range monthValues {
		fmt.Printf("Month: %s, Value: %d\n", mv.Month, mv.Value)
	}
}

type monthValue struct {
	Month string `json:"month"`
	Value int    `json:"value"`
}

func GetMonthValues(mode string) []monthValue {
	var months []monthValue

	currentTime := time.Now()

	switch mode {
	case "current":
		for i := 5; i >= 0; i-- {
			month := currentTime.AddDate(0, -i, 0).Format("January")
			months = append(months, monthValue{
				Month: month,
			})
		}
	case "halfleft":
		for i := 0; i < 6; i++ {
			month := time.Month(i + 1).String()
			months = append(months, monthValue{
				Month: month,
			})
		}
	case "halfright":
		for i := 6; i < 12; i++ {
			month := time.Month(i + 1).String()
			months = append(months, monthValue{
				Month: month,
			})
		}
	}

	return months
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	res := 1
	for _, num := range nums {
		if num > 0 {
			res *= 1
		} else if num < 0 {
			res *= -1
		} else if num == 0 {
			return 0
		}
	}

	return res
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sSplit := strings.Split(s, "")
	tSplit := strings.Split(t, "")

	sort.Strings(sSplit)
	sort.Strings(tSplit)

	for i := 0; i < len(sSplit); i++ {
		fmt.Println(tSplit[i], sSplit[i])
		if sSplit[i] != tSplit[i] {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	var diff byte
	i := 0
	for i = 0; i < len(s); i++ {
		diff += t[i]
		diff -= s[i]
	}

	if i < len(t) {
		return diff + t[i]
	} else {
		return diff
	}
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr)
	if n == 2 {
		return true
	}
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	const (
		TOTAL_CARDS = 52
		SYMBOL      = 4
		NUMBER      = 13
	)

	d.cards = make([]Card, 0, TOTAL_CARDS)
	for i := 0; i < SYMBOL; i++ {
		for j := 1; j <= NUMBER; j++ {
			d.cards = append(d.cards, Card{symbol: i, number: j})
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
