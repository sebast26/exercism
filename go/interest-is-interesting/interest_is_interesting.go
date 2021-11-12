package interest

const negativeBalance float32 = 3.213
const positiveBalance float32 = 0.5
const smallPositiveBalance float32 = 1.621
const bigPositiveBalance float32 = 2.475

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return negativeBalance
	case balance < 1000:
		return positiveBalance
	case balance >= 1000 && balance < 5000:
		return smallPositiveBalance
	default:
		return bigPositiveBalance
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}
	return years
}
