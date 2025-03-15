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
	_ = []Card{
		{ID: 1, Symbols: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{ID: 2, Symbols: []int{1, 9, 10, 11, 12, 13, 14, 15}},
		{ID: 3, Symbols: []int{1, 16, 17, 18, 19, 20, 21, 22}},
		{ID: 4, Symbols: []int{1, 23, 24, 25, 26, 27, 28, 29}},
		{ID: 5, Symbols: []int{1, 30, 31, 32, 33, 34, 35, 36}},
		{ID: 6, Symbols: []int{1, 37, 38, 39, 40, 41, 42, 43}},
		{ID: 7, Symbols: []int{1, 44, 45, 46, 47, 48, 49, 50}},
		{ID: 8, Symbols: []int{1, 51, 52, 53, 54, 55, 56, 57}},
		//
		{ID: 9, Symbols: []int{2, 9, 16, 23, 30, 37, 44, 51}},
		{ID: 10, Symbols: []int{2, 10, 17, 24, 31, 38, 45, 52}},
		{ID: 11, Symbols: []int{2, 11, 18, 25, 32, 39, 46, 53}},
		{ID: 12, Symbols: []int{2, 12, 19, 26, 33, 40, 47, 54}},
		{ID: 13, Symbols: []int{2, 13, 20, 27, 34, 41, 48, 55}},
		{ID: 14, Symbols: []int{2, 14, 21, 28, 35, 42, 49, 56}},
		{ID: 15, Symbols: []int{2, 15, 22, 29, 36, 43, 50, 57}},
		//
		{ID: 16, Symbols: []int{3, 9, 17, 25, 33, 41, 49, 57}},
		{ID: 17, Symbols: []int{3, 10, 18, 26, 34, 42, 50, 51}},
		{ID: 18, Symbols: []int{3, 11, 19, 27, 35, 43, 44, 52}},
		{ID: 19, Symbols: []int{3, 12, 20, 28, 36, 37, 45, 53}},
		{ID: 20, Symbols: []int{3, 13, 21, 29, 30, 38, 46, 54}},
		{ID: 21, Symbols: []int{3, 14, 22, 23, 31, 39, 47, 55}},
		{ID: 22, Symbols: []int{3, 15, 16, 24, 32, 40, 48, 56}},
		//
		{ID: 23, Symbols: []int{4, 9, 18, 27, 36, 38, 47, 56}},
		{ID: 24, Symbols: []int{4, 10, 19, 28, 30, 39, 48, 57}},
		{ID: 25, Symbols: []int{4}},
		{ID: 26, Symbols: []int{4}},
		{ID: 27, Symbols: []int{4}},
		{ID: 28, Symbols: []int{4}},
		{ID: 29, Symbols: []int{4}},
		//
		{ID: 30, Symbols: []int{5, 9, 19, 29, 32, 42, 45, 55}},
		{ID: 31, Symbols: []int{5}},
		{ID: 32, Symbols: []int{5}},
		{ID: 33, Symbols: []int{5}},
		{ID: 34, Symbols: []int{5}},
		{ID: 35, Symbols: []int{5}},
		{ID: 36, Symbols: []int{5}},
		//
		{ID: 37, Symbols: []int{6, 9, 20, 24}},
		{ID: 38, Symbols: []int{6}},
		{ID: 39, Symbols: []int{6}},
		{ID: 40, Symbols: []int{6}},
		{ID: 41, Symbols: []int{6}},
		{ID: 42, Symbols: []int{6}},
		{ID: 43, Symbols: []int{6}},
		//
		{ID: 44, Symbols: []int{7, 9, 21}},
		{ID: 45, Symbols: []int{7}},
		{ID: 46, Symbols: []int{7}},
		{ID: 47, Symbols: []int{7}},
		{ID: 48, Symbols: []int{7}},
		{ID: 49, Symbols: []int{7}},
		{ID: 50, Symbols: []int{7}},
		//
		{ID: 51, Symbols: []int{8, 9, 22}},
		{ID: 52, Symbols: []int{8}},
		{ID: 53, Symbols: []int{8}},
		{ID: 54, Symbols: []int{8}},
		{ID: 55, Symbols: []int{8}},
		{ID: 56, Symbols: []int{8}},
		{ID: 57, Symbols: []int{8}},
	}

	//Example cards
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
	if numberOfMatches != 1 {
		fmt.Println("mismatch")
		fmt.Println(newCard)
		fmt.Println(existingCard)
	}
	return numberOfMatches == 1
}
