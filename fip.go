package main

type FipResponse struct {
	Data struct {
		Now struct {
			PlayingItem struct {
				Title     string `json:"title"`
				Subtitle  string `json:"subtitle"`
				Cover     string `json:"cover"`
				StartTime int64  `json:"start_time"`
				EndTime   int64  `json:"end_time"`
			} `json:"playing_item"`
			Program struct {
				UUID          string `json:"uuid"`
				Name          string `json:"name"`
				VisualConcept struct {
					URL     string `json:"url"`
					Preview string `json:"preview"`
					Width   int    `json:"width"`
					Height  int    `json:"height"`
				} `json:"visualConcept"`
			} `json:"program"`
			Song struct {
				Cover         string        `json:"cover"`
				Title         string        `json:"title"`
				Interpreters  []string      `json:"interpreters"`
				MusicalKind   string        `json:"musical_kind"`
				Label         string        `json:"label"`
				Album         string        `json:"album"`
				ExternalLinks ExternalLinks `json:"external_links"`
			} `json:"song"`
			ServerTime  int64  `json:"server_time"`
			NextRefresh int64  `json:"next_refresh"`
			Mode        string `json:"mode"`
		} `json:"now"`
		PreviousTracks struct {
			Edges []struct {
				Node struct {
					Title     string `json:"title"`
					Subtitle  string `json:"subtitle"`
					StartTime int64  `json:"start_time"`
					Cover     string `json:"cover"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"previousTracks"`
		NextTracks []struct {
			Title     string `json:"title"`
			Subtitle  string `json:"subtitle"`
			StartTime int64  `json:"start_time"`
			Cover     string `json:"cover"`
		} `json:"nextTracks"`
	} `json:"data"`
}

type ExternalLinks struct {
	YouTube string `json:"youtube"`
	Deezer  struct {
		ID    string `json:"id"`
		Link  string `json:"link"`
		Image string `json:"image"`
	} `json:"deezer"`
	Itunes struct {
		ID    string `json:"id"`
		Link  string `json:"link"`
		Image string `json:"image"`
	} `json:"itunes"`
	Spotify struct {
		ID    string `json:"id"`
		Link  string `json:"link"`
		Image string `json:"image"`
	} `json:"spotify"`
}
