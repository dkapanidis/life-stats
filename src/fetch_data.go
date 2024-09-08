package main

import (
	"encoding/json"
	"fmt"
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

	// body, _ := ioutil.ReadAll(resp.Body)

	// var activities []StravaActivity
	// json.Unmarshal(body, &activities)
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
	os.WriteFile(filename, prettyJSON, 0644)
	fmt.Println("Strava data saved to", filename)
}

func main() {
	stravaAccessToken := os.Getenv("STRAVA_ACCESS_TOKEN")

	fetchStravaData(stravaAccessToken)
}
