package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValues := map[string]int{
        "ace": 11, "king": 10, "queen": 10, "jack": 10,
        "ten": 10, "nine": 9, "eight": 8, "seven": 7, 
        "six": 6, "five": 5, "four": 4, "three": 3, "two": 2,
    }
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myHand := ParseCard(card1) + ParseCard(card2)
    switch myHand {
    case 22:
    	return "P"
    case 21:
    	switch {
        case ParseCard(dealerCard) >= 10:
        	return "S"
        default:
        	return "W"
        }
    case 17, 18, 19, 20:
    	return "S"
    case 12, 13, 14, 15, 16:
    	switch {
        case ParseCard(dealerCard) >= 7:
        	return "H"
        default:
        	return "S"
        }
    default:
    	return "H"
    }
}
