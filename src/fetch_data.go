package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Structs to parse the JSON response
type StravaActivity struct {
	// Add fields relevant to your Strava data
	Name string `json:"name"`
	// Add other fields as necessary
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

	body, _ := ioutil.ReadAll(resp.Body)

	var activities []StravaActivity
	json.Unmarshal(body, &activities)

	filename := fmt.Sprintf("strava_data_%s.json", time.Now().Format("20060102"))
	ioutil.WriteFile(filename, body, 0644)
	fmt.Println("Strava data saved to", filename)
}

func main() {
	stravaAccessToken := os.Getenv("STRAVA_ACCESS_TOKEN")

	fetchStravaData(stravaAccessToken)
}
