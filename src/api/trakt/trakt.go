package trakt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dkapanidis/life-stats/src/lib/storage"
	"github.com/dkapanidis/life-stats/src/models"
	"github.com/sirupsen/logrus"
)

func refreshTraktToken(clientID, clientSecret, refreshToken string) (string, error) {
	url := "https://api.trakt.tv/oauth/token"
	body := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
		"redirect_uri":  "urn:ietf:wg:oauth:2.0:oob",
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %v", err)
	}
	defer resp.Body.Close()

	var tokenResp interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %v", err)
	}
	logrus.Info(tokenResp)
	return "", nil
}

func fetchTraktData() {
	resp, err := fetch("/sync/watched/shows")
	if err != nil {
		fmt.Println("Failed to fetch Trakt data:", err)
		return
	}
	defer resp.Body.Close()

	var data = make([]models.WatchedItem, 0)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Failed to parse Trakt JSON:", err)
		return
	}

	storage.StoreTo(data, "data/trakt/api.json")

	ratings, err := fetchTraktRatings()
	if err != nil {
		fmt.Println("Failed to fetch trakt ratings:", err)
		return
	}

	storage.StoreTo(ToShows(data, ratings), "data/trakt/summary.json")
}

func fetchTraktRatings() ([]models.RatingItem, error) {
	resp, err := fetch("/users/me/ratings/shows")
	if err != nil {
		fmt.Println("Failed to fetch Trakt data:", err)
		return []models.RatingItem{}, err
	}
	defer resp.Body.Close()

	var data = make([]models.RatingItem, 0)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Failed to parse Trakt JSON:", err)
		return []models.RatingItem{}, err
	}

	storage.StoreTo(data, "data/trakt/ratings.json")
	return data, nil
}

func fetch(url string) (*http.Response, error) {
	accessToken := os.Getenv("TRAKT_ACCESS_TOKEN")

	req, _ := http.NewRequest("GET", "https://api.trakt.tv"+url, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", os.Getenv("TRAKT_CLIENT_ID"))
	req.Header.Add("Content-type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func Sync() {
	fetchTraktData()
}
