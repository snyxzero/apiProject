package service

type СalculationRatingPoints struct {
}

func NewRatingPoints() *СalculationRatingPoints {
	return &СalculationRatingPoints{}
}

func (o *СalculationRatingPoints) CalculateRatingPointsToUser(numberOfRatingsFromUser int, numberOfRatingsForBreweryFromUser int) int {
	points := 0

	if numberOfRatingsFromUser == 1 {
		points += 50
	}

	if numberOfRatingsFromUser%3 == 0 {
		points += 5
	}

	if numberOfRatingsForBreweryFromUser%2 == 0 {
		points += 10
	}

	return points
}
