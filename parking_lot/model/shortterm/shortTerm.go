package shortterm

const (
	hourInMinutes int = 60
	dayInMinutes  int = 24 * hourInMinutes
)

// ShortTermCalculator represents the Short Term parking modality
type ShortTermCalculator struct {
	minimumCost                    float32
	costToBeAdded                  float32
	costToBeAddedIntervalInMinutes int
	dailyMaximumCost               float32
}

// New returns an instance of the ShortTermCalculator object
func New() *ShortTermCalculator {
	return &ShortTermCalculator{
		minimumCost:                    2.00,
		costToBeAdded:                  1.00,
		costToBeAddedIntervalInMinutes: 30,
		dailyMaximumCost:               24.00,
	}
}

// CalculateParkingCost effectively calculates the value to be charged given a specific duration time
func (s *ShortTermCalculator) CalculateParkingCost(minutesSpent int) float32 {
	var overallCost float32

	if minutesSpent > dayInMinutes {
		numberOfDays, leftover := getNumberOfDaysSpentWithMinutesLeftovers(minutesSpent)
		overallCost = float32(numberOfDays) * s.dailyMaximumCost
		minutesSpent = leftover
	}

	partialCost := float32(getCostMultiplier(minutesSpent, s.costToBeAddedIntervalInMinutes)) * s.costToBeAdded

	if partialCost <= s.dailyMaximumCost {
		overallCost += partialCost
	} else {
		return s.dailyMaximumCost
	}

	if overallCost < s.minimumCost {
		return s.minimumCost
	}

	return overallCost
}

func getNumberOfDaysSpentWithMinutesLeftovers(minutesSpent int) (int, int) {
	if minutesSpent > 0 {
		daysSpent := minutesSpent / dayInMinutes
		leftoverTime := minutesSpent % dayInMinutes
		return daysSpent, leftoverTime
	}
	return 0, 0
}

func getCostMultiplier(minutesSpent int, costInvervalInMinutes int) int {
	if minutesSpent > 0 {
		multiplier := minutesSpent / costInvervalInMinutes
		leftover := minutesSpent % costInvervalInMinutes
		if leftover > 0 {
			multiplier++
		}
		return multiplier
	}
	return 0
}
