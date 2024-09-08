package trakt

import (
	"github.com/dkapanidis/life-stats/src/api/fanart"
	"github.com/dkapanidis/life-stats/src/models"
	"github.com/sirupsen/logrus"
)

func ToShows(items []models.WatchedItem, ratings []models.RatingItem) []models.Show {
	response := make([]models.Show, 0)
	for _, item := range items {
		var rating *int
		for _, ratingItem := range ratings {
			if ratingItem.Show.IDs.Slug == item.Show.IDs.Slug {
				rating = &ratingItem.Rating
			}
		}
		response = append(response, ToShow(item, rating))
	}
	return response
}

func ToShow(item models.WatchedItem, rating *int) models.Show {
	url, err := fanart.FetchFanartThumbnail(item.Show.IDs.TVDB)
	if err != nil {
		logrus.Warnf("Cannot get fanart thumbnail for %s: %s", item.Show.Title, err)
	}
	return models.Show{
		Title:  item.Show.Title,
		URL:    url,
		Rating: rating,
	}
}
