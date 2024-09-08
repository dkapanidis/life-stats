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

func RefreshTraktToken(clientID, clientSecret, refreshToken string) (string, error) {
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

func FetchTraktData() {
	accessToken := os.Getenv("TRAKT_ACCESS_TOKEN")

	url := "https://api.trakt.tv/sync/watched/shows" // Adjust this endpoint based on what you need
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", os.Getenv("TRAKT_CLIENT_ID"))
	req.Header.Add("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to fetch Trakt data:", err)
		return
	}
	defer resp.Body.Close()
	logrus.Info(resp.Status)
	var data = make([]models.WatchedItem, 0)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Failed to parse Trakt JSON:", err)
		return
	}

	storage.StoreTo(data, "data/trakt/api.json")
}
