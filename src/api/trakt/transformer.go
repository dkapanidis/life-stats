package trakt

import "github.com/dkapanidis/life-stats/src/models"

func ToShows(items []models.WatchedItem) []models.Show {
	response := make([]models.Show, 0)
	for _, item := range items {
		response = append(response, ToShow(item))
	}
	return response
}

func ToShow(item models.WatchedItem) models.Show {
	return models.Show{
		Title: item.Show.Title,
	}
}
