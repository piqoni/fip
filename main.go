package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Make the HTTP request to fetch the JSON data
	response, err := http.Get("https://www.fip.fr/latest/api/graphql?operationName=Now&variables={%22bannerPreset%22%3A%22600x600-noTransform%22%2C%22stationId%22%3A7%2C%22previousTrackLimit%22%3A3}&extensions={%22persistedQuery%22%3A{%22version%22%3A1%2C%22sha256Hash%22%3A%228a931c7d177ff69709a79f4c213bd2403f0c11836c560bc22da55628d8100df8%22}}")
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}

	// Create an instance of FipResponse to hold the unmarshaled data
	fipResponse := FipResponse{}

	// Unmarshal the JSON response into the FipResponse struct
	err = json.Unmarshal([]byte(responseBody), &fipResponse)
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	next := flag.Bool("next", false, "Specify --next to show the next track")
	now := flag.Bool("now", false, "Specify --now to show the current track")
	prev := flag.Bool("prev", false, "Specify --prev to show the previous track")
	spotify := flag.Bool("spotify", false, "Specify --spotify to add spotify link at the end of the track")
	flag.Parse()

	var spotifyLink string

	if *spotify && fipResponse.Data.Now.Song.ExternalLinks.Spotify.ID != "" {
		spotifyLink = fipResponse.Data.Now.Song.ExternalLinks.Spotify.Link
	} else {
		spotifyLink = ""
	}

	prevPlaying, nextPlaying, nowPlaying := "", "", ""

	if fipResponse.Data.Now.PlayingItem.Title != "" {
		nowPlaying = fipResponse.Data.Now.PlayingItem.Title + " - " + fipResponse.Data.Now.PlayingItem.Subtitle + " " + spotifyLink
	}

	if len(fipResponse.Data.PreviousTracks.Edges) > 0 {
		prevPlaying = fipResponse.Data.PreviousTracks.Edges[0].Node.Title + " - " + fipResponse.Data.PreviousTracks.Edges[0].Node.Subtitle + " " + spotifyLink
	}

	if len(fipResponse.Data.NextTracks) > 0 {
		nextPlaying = fipResponse.Data.NextTracks[0].Title + " - " + fipResponse.Data.NextTracks[0].Subtitle + " " + spotifyLink
	}

	if *next {
		fmt.Println(nextPlaying)
	}
	// If --now or no option is specified (default behavior), print the current track
	if *now || (!*next && !*prev) {
		fmt.Println(nowPlaying)
	}
	if *prev {
		fmt.Println(prevPlaying)
	}
}
