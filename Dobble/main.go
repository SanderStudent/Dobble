package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Card struct {
	ID      int
	Symbols []int
}

// 57 symbols, can make up to 57 cards with 8 symbols each. Each card has exactly 1 matching symbol with each other card.
func main() {
	//Example cards for 3 symbols per card
	_ = []Card{
		{ID: 1, Symbols: []int{1, 2, 3}},
		{ID: 2, Symbols: []int{1, 4, 5}},
		{ID: 3, Symbols: []int{1, 6, 7}},
		{ID: 4, Symbols: []int{2, 4, 6}},
		{ID: 5, Symbols: []int{2, 5, 7}},
		{ID: 6, Symbols: []int{3, 4, 7}},
		{ID: 7, Symbols: []int{3, 5, 6}},
	}
	//Choose the number of cards that need to be visible on each card
	numberOfSymbolsPerCard := 8
	generatedCards := generateCards(numberOfSymbolsPerCard)
	for _, card := range generatedCards {
		fmt.Println(card)
	}
	fmt.Println(checkGeneratedDeck(generatedCards)) //Checks if the generated deck is valid
}

func checkGeneratedDeck(cards []Card) bool {
	for i := range cards {
		checkCard := cards[i]
		firstPartOfDeck := cards[:i]
		secondPartOfDeck := cards[i+1:]
		if !isCardValid(checkCard, firstPartOfDeck) || !isCardValid(checkCard, secondPartOfDeck) {
			return false
		}
	}
	return true
}

func matchesOnce(newCard Card, existingCard Card) bool {
	if len(newCard.Symbols) != len(existingCard.Symbols) {
		return false
	}
	numberOfMatches := 0
	for _, symbol := range newCard.Symbols {
		if slices.Contains(existingCard.Symbols, symbol) {
			numberOfMatches++
		}
	}
	return numberOfMatches == 1
}

func generateCards(numberOfSymbolsOnCard int) []Card {
	baseCard := Card{ID: 1}
	firstSymbol := 1
	firstSymbolUsed := 1
	k := 1
	m := 1
	n := 1
	shift := 0
	numberOfCardsWithCurrentShift := 0
	for i := 1; i <= numberOfSymbolsOnCard; i++ {
		baseCard.Symbols = append(baseCard.Symbols, i)
		k++
	}
	result := []Card{baseCard}
	numberOfCards := 1 + numberOfSymbolsOnCard*(numberOfSymbolsOnCard-1)

	for i := 2; i <= numberOfCards; i++ {
		newCard := Card{ID: i}
		for j := 1; j <= numberOfSymbolsOnCard; j++ {
			if j == 1 {
				newCard.Symbols = append(newCard.Symbols, firstSymbol)
				firstSymbolUsed++
				if firstSymbolUsed == numberOfSymbolsOnCard {
					firstSymbol++
					firstSymbolUsed = 1
				}
			} else if i <= numberOfSymbolsOnCard {
				newCard.Symbols = append(newCard.Symbols, k)
				k++
			} else {
				k = result[m].Symbols[n]
				m++
				n += shift
				n = 1 + ((n - 1) % (numberOfSymbolsOnCard - 1))
				newCard.Symbols = append(newCard.Symbols, k)
				if j == numberOfSymbolsOnCard { //end of card
					m = 1
					n++
					n = 1 + ((n - 1) % (numberOfSymbolsOnCard - 1))
					numberOfCardsWithCurrentShift++
					if numberOfCardsWithCurrentShift == numberOfSymbolsOnCard-1 {
						shift++
						numberOfCardsWithCurrentShift = 0
					}
				}
			}
		}
		result = append(result, newCard)
	}
	return result
}

func isCardValid(card Card, existingCards []Card) bool {
	for _, existingCard := range existingCards {
		if !matchesOnce(card, existingCard) {

			return false
		}
	}
	return true
}
