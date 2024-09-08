package strava

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dkapanidis/life-stats/src/lib/storage"
	"github.com/dkapanidis/life-stats/src/models"
)

func refreshStravaToken(clientID, clientSecret, refreshToken string) (string, error) {
	url := "https://www.strava.com/api/v3/oauth/token"
	body := fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s",
		clientID, clientSecret, refreshToken,
	)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %v", err)
	}
	defer resp.Body.Close()

	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %v", err)
	}

	return tokenResp.AccessToken, nil
}

func fetchStravaData(accessToken string) {
	url := "https://www.strava.com/api/v3/athlete/activities"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to fetch Strava data:", err)
		return
	}
	defer resp.Body.Close()

	var data = make([]models.StravaActivity, 0)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Failed to parse Strava JSON:", err)
		return
	}

	storage.StoreTo(data, "data/strava/api.json")
	storage.StoreTo(ToRunnings(data), "data/strava/summary.json")
}

func Sync() {
	// Variables
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	refreshToken := os.Getenv("STRAVA_REFRESH_TOKEN")

	// Refresh the access token
	accessToken, err := refreshStravaToken(clientID, clientSecret, refreshToken)
	if err != nil {
		fmt.Println("Error refreshing Strava token:", err)
		return
	}

	// Fetch data
	fetchStravaData(accessToken)
}
