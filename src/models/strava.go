package models

type StravaActivity struct {
	AchivementCount            int64         `json:"achievement_count"`
	Athlete                    StravaAthlete `json:"athlete"`
	AthleteCount               int64         `json:"athlete_count"`
	AverageSpeed               float64       `json:"average_speed"`
	CommentCount               int64         `json:"comment_count"`
	Commute                    bool          `json:"commute"`
	DisplayHideHeartrateOption bool          `json:"display_hide_heartrate_option"`
	Distance                   float64       `json:"distance"`
	ElapsedTime                int64         `json:"elapsed_time"`
	EndLatlng                  []float64     `json:"end_latlng"`
	ExternalID                 string        `json:"external_id"`
	Flagged                    bool          `json:"flagged"`
	FromAcceptedTag            bool          `json:"from_accepted_tag"`
	GearID                     string        `json:"gear_id"`
	HasHeartrate               bool          `json:"has_heartrate"`
	HasKudoed                  bool          `json:"has_kudoed"`
	HeartrateOptOut            bool          `json:"heartrate_opt_out"`
	ID                         int64         `json:"id"`
	KudosCount                 int64         `json:"kudos_count"`
	LocationCity               string        `json:"location_city"`
	LocationCountry            string        `json:"location_country"`
	LocationState              string        `json:"location_state"`
	Manual                     bool          `json:"manual"`
	Map                        StravaMap     `json:"map"`
	MaxSpeed                   float64       `json:"max_speed"`
	MovingTime                 int64         `json:"moving_time"`
	Name                       string        `json:"name"`
	PhotoCount                 int64         `json:"photo_count"`
	PRCount                    int64         `json:"pr_count"`
	Private                    bool          `json:"private"`
	ResourceState              int64         `json:"resource_state"`
	SportType                  string        `json:"sport_type"`
	StartDate                  string        `json:"start_date"`
	StartDateLocal             string        `json:"start_date_local"`
	StartLatlng                []float64     `json:"start_latlng"`
	Timezone                   string        `json:"timezone"`
	TotalElevationGain         float64       `json:"total_elevation_gain"`
	TotalPhotoCount            int64         `json:"total_photo_count"`
	Trainer                    bool          `json:"trainer"`
	Type                       string        `json:"type"`
	UploadID                   int64         `json:"upload_id"`
	UTCOffset                  float64       `json:"utc_offset"`
	Visibility                 string        `json:"visibility"`
	WorkoutType                int64         `json:"workout_type"`
}

type StravaAthlete struct {
	ID            int64 `json:"id"`
	ResourceState int64 `json:"resource_state"`
}

type StravaMap struct {
	ID              string `json:"id"`
	ResourceState   int64  `json:"resource_state"`
	SummaryPolyline string `json:"summary_polyline"`
}
