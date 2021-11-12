package blackjack

const BlackJack int = 21
const Stand string = "S"
const Hit string = "H"
const Split string = "P"
const AutomaticallyWin = "W"

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

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == BlackJack
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	dealerAce := dealerScore == ParseCard("ace")
	dealerFigure := dealerScore == ParseCard("jack") ||
		dealerScore == ParseCard("queen") ||
		dealerScore == ParseCard("king")
	dealerTen := dealerScore == ParseCard("ten")
	if !isBlackjack {
		return Split
	}
	if !dealerAce && !dealerFigure && !dealerTen {
		return AutomaticallyWin
	}
	return Stand
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return Stand
	}
	if handScore <= 11 || dealerScore >= 7 {
		return Hit
	}
	return Stand
}

// FirstTurn returns the semi-optimal decision for the first turn, given the cards of the player and the dealer.
// This function is already implemented and does not need to be edited. It pulls the other functions together in a
// complete decision tree for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if 20 < handScore {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}
	return SmallHand(handScore, dealerScore)
}
