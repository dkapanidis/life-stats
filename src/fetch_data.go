package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Structs for token handling
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

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

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %v", err)
	}

	// Update the refresh token and access token in environment variables or secrets if needed.
	// fmt.Println("New Access Token:", tokenResp.AccessToken)
	// fmt.Println("New Refresh Token:", tokenResp.RefreshToken)

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

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Failed to parse Strava JSON:", err)
		return
	}

	// Pretty-print the JSON data
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Failed to format Strava data:", err)
		return
	}

	filename := fmt.Sprintf("strava_data_%s.json", time.Now().Format("20060102"))
	ioutil.WriteFile(filename, prettyJSON, 0644)
	fmt.Println("Strava data saved to", filename)
}

func main() {
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	refreshToken := os.Getenv("STRAVA_REFRESH_TOKEN")

	// Refresh the access token
	accessToken, err := refreshStravaToken(clientID, clientSecret, refreshToken)
	if err != nil {
		fmt.Println("Error refreshing Strava token:", err)
		return
	}

	// Fetch data using the refreshed access token
	fetchStravaData(accessToken)
}
