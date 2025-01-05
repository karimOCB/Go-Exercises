package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
        case "two":
        	return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
        case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
        	return 9
        case "ten":
        	return 10
        case "jack":
        	return 10
        case "queen":
        	return 10
        case "king":
        	return 10
        default:
        	return 0
    
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1N := ParseCard(card1)
    card2N := ParseCard(card2)
    cardSum := card1N + card2N
    card3N := ParseCard(dealerCard)
    
	if card1N == 11 && card2N == 11 {
        return "P"
    }
    
    if cardSum == 21 {
        if card3N != 11 && card3N != 10 {
            return "W"
        }
        return "S"
    }
    
    if cardSum >= 17 && cardSum <= 20 {
        return "S"
    } else if cardSum >= 12 && cardSum <= 16 {
        if card3N >= 7 {
            return "H"
        } 
        return "S"
    } else {
        return "H"
    }
}
