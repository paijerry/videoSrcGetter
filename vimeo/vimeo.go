package vimeo

import (
	"test/videoSrcGetter/env"

	"encoding/xml"

	"github.com/paijerry/ezapi"
)

// RspnVideoInfo -
type RspnVideoInfo struct {
	XMLName xml.Name  `xml:"videos"`
	Vedio   VideoInfo `xml:"video"`
}

// VideoInfo -
type VideoInfo struct {
	XMLName       xml.Name `xml:"video"  json:"video"`
	ID            string   `xml:"id"  json:"id"`
	Title         string   `xml:"title"  json:"title"`
	Description   string   `xml:"description"  json:"description"`
	URL           string   `xml:"url"  json:"url"`
	ImageS        string   `xml:"thumbnail_small"  json:"thumbnail_small"`
	ImageM        string   `xml:"thumbnail_medium"  json:"thumbnail_medium"`
	ImageL        string   `xml:"thumbnail_large" json:"thumbnail_large"`
	UploadDate    string   `xml:"upload_date" json:"upload_date"`
	MobileURL     string   `xml:"mobile_url" json:"mobile_url"`
	UserID        string   `xml:"user_id" json:"user_id"`
	UserName      string   `xml:"user_name" json:"user_name"`
	UserURL       string   `xml:"user_url" json:"user_url"`
	UserPortraitS string   `xml:"user_portrait_small" json:"user_portrait_small"`
	UserPortraitM string   `xml:"user_portrait_medium" json:"user_portrait_medium"`
	UserPortraitL string   `xml:"user_portrait_large" json:"user_portrait_large"`
	UserPortraitH string   `xml:"user_portrait_huge" json:"user_portrait_huge"`
	Duration      string   `xml:"duration" json:"duration"`
	Width         string   `xml:"width" json:"width"`
	Height        string   `xml:"height" json:"height"`
	Tags          string   `xml:"tags" json:"tags"`
	EmbedPrivacy  string   `xml:"embed_privacy" json:"embed_privacy"`
}

// GetInfo -
func GetInfo(vid string) (videoInfo VideoInfo, err error) {
	rspn, err := ezapi.New().URL(env.VimeoAPI + vid + ".xml").Do("GET")
	if err != nil {
		return
	}

	result := new(RspnVideoInfo)

	xml.Unmarshal(rspn.Body, &result)
	videoInfo = result.Vedio
	return
}
