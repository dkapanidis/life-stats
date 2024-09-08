package trakt

import (
	"github.com/dkapanidis/life-stats/src/api/fanart"
	"github.com/dkapanidis/life-stats/src/models"
	"github.com/sirupsen/logrus"
)

func ToShows(items []models.WatchedItem) []models.Show {
	response := make([]models.Show, 0)
	for _, item := range items {
		response = append(response, ToShow(item))
	}
	return response
}

func ToShow(item models.WatchedItem) models.Show {
	url, err := fanart.FetchFanartThumbnail(item.Show.IDs.TVDB)
	if err != nil {
		logrus.Warnf("Cannot get fanart thumbnail for %s: %s", item.Show.Title, err)
	}
	return models.Show{
		Title: item.Show.Title,
		URL:   url,
	}
}
