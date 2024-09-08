package fanart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Struct for Fanart.tv API response to get images
type FanartImages struct {
	Name           string        `json:"name"`
	TheTVDB_ID     string        `json:"thetvdb_id"`
	HDTVLogo       []FanartImage `json:"hdtvlogo"`
	TVPoster       []FanartImage `json:"tvposter"`
	ShowBackground []FanartImage `json:"showbackground"`
	HDClearArt     []FanartImage `json:"hdclearart"`
	ClearLogo      []FanartImage `json:"clearlogo"`
	SeasonPoster   []FanartImage `json:"seasonposter"`
	TVThumb        []FanartImage `json:"tvthumb"`
	SeasonThumb    []FanartImage `json:"seasonthumb"`
	CharacterArt   []FanartImage `json:"characterart"`
	TVBanner       []FanartImage `json:"tvbanner"`
	ClearArt       []FanartImage `json:"clearart"`
	SeasonBanner   []FanartImage `json:"seasonbanner"`
}
type FanartImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Lang   string `json:"lang"`
	Likes  string `json:"likes"`
	Season string `json:"season,omitempty"`
}

type TVShowImages struct {
	Images []ImageDetails `json:"images"`
}

type ImageDetails struct {
	URL string `json:"url"`
}

func FetchFanartThumbnail(tmdbID int) (string, error) {
	fanartAPIKey := os.Getenv("FANART_API_KEY")
	client := &http.Client{}

	// Fetch images from Fanart.tv
	fanartURL := fmt.Sprintf("https://webservice.fanart.tv/v3/tv/%d?api_key=%s", tmdbID, fanartAPIKey)
	req, err := http.NewRequest("GET", fanartURL, nil)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create Fanart.tv request:")
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "Failed to fetch Fanart.tv data:")
	}
	defer resp.Body.Close()

	var fanartImages FanartImages
	if err := json.NewDecoder(resp.Body).Decode(&fanartImages); err != nil {
		return "", errors.Wrap(err, "Failed to parse Fanart.tv JSON:")
	}

	if len(fanartImages.TVThumb) > 0 {
		thumbnailURL := fanartImages.TVThumb[0].URL
		previewURL := strings.Replace(thumbnailURL, "assets.fanart.tv/fanart", "assets.fanart.tv/preview", 1)
		return previewURL, nil
	} else {
		return "", errors.New("No thumbnail available for this show.")
	}

}
