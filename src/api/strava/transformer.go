package strava

import "github.com/dkapanidis/life-stats/src/models"

func ToRunnings(activities []models.StravaActivity) []models.Running {
	response := make([]models.Running, 0)
	for _, activity := range activities {
		response = append(response, ToRunning(activity))
	}
	return response
}

func ToRunning(activity models.StravaActivity) models.Running {
	return models.Running{
		StartDate: activity.StartDate,
		Distance:  activity.Distance,
	}
}
