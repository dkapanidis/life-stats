package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func StoreTo(data interface{}, filename string) {
	// Pretty-print the JSON data
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Failed to format Strava data:", err)
		return
	}

	os.WriteFile(filename, prettyJSON, 0644)
	fmt.Println("Strava data saved to", filename)

}
