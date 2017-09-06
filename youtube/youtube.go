package youtube

import (
	"encoding/json"
	"test/videoSrcGetter/env"

	"github.com/paijerry/ezapi"
)

// RspnVideoInfo -
type RspnVideoInfo struct {
	HTML            string `json:"html"`
	ThumbnailWidth  int64  `json:"thumbnail_width"`
	ThumbnailHeight int64  `json:"thumbnail_height"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailURL1   string `json:"thumbnail_url_1"`
	ThumbnailURL2   string `json:"thumbnail_url_2"`
	ThumbnailURL3   string `json:"thumbnail_url_3"`
	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
	Width           int64  `json:"width"`
	Height          int64  `json:"height"`
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	Type            string `json:"type"`
	Version         string `json:"version"`
	Title           string `json:"title"`
}

// GetInfo -
func GetInfo(vid string) (result RspnVideoInfo, err error) {
	rspn, err := ezapi.New().URL(env.YouTubeAPI + vid).Do("GET")
	if err != nil {
		return
	}

	json.Unmarshal(rspn.Body, &result)

	result.ThumbnailURL1 = env.YouTubeImgURL + vid + "/hq1.jpg"
	result.ThumbnailURL2 = env.YouTubeImgURL + vid + "/hq2.jpg"
	result.ThumbnailURL3 = env.YouTubeImgURL + vid + "/hq3.jpg"
	return
}
