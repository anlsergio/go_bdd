package valet

const (
	hourInMinutes int = 60
	dayInMinutes  int = 24 * hourInMinutes
)

// ValetCalculator represents the Valet parking modality
type ValetCalculator struct {
	dailyCost                  float32
	discount                   float32
	discountThresholdInMinutes int
}

// New returns an instance of the ValetCalculator object
func New() *ValetCalculator {
	return &ValetCalculator{
		dailyCost:                  18.00,
		discount:                   6.00,
		discountThresholdInMinutes: 5 * hourInMinutes,
	}
}

// CalculateParkingCost effectively calculates the value to be charged given a specific duration time
func (v *ValetCalculator) CalculateParkingCost(minutesSpent int) float32 {
	if minutesSpent <= v.discountThresholdInMinutes {
		return v.dailyCost - v.discount
	}

	return float32(getNumberOfDaysSpent(minutesSpent)) * v.dailyCost
}

func getNumberOfDaysSpent(minutesSpent int) int {
	if minutesSpent > 0 {
		daysSpent := minutesSpent / dayInMinutes
		leftoverTime := minutesSpent % dayInMinutes
		if leftoverTime > 0 {
			daysSpent++
		}
		return daysSpent
	}
	return 0
}
