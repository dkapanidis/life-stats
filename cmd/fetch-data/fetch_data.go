package main

import (
	"github.com/dkapanidis/life-stats/src/api/strava"
	"github.com/dkapanidis/life-stats/src/api/trakt"
)

func main() {
	strava.Sync()
	trakt.Sync()
}
