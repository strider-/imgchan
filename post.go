package imgchan

import (
	"encoding/base64"
	"fmt"
)

type Post struct {
	Board           string // Not part of the response, but appended for library usage
	Number          int    `json:"no"`
	Name            string
	MD5             string
	Subject         string `json:"sub"`
	Comment         string `json:"com"`
	Replies         int
	Images          int
	Filename        string
	Ext             string
	Filesize        int `json:"fsize"`
	Time            int64
	ImageWidth      int   `json:"w"`
	ImageHeight     int   `json:"h"`
	ThumbnailWidth  int   `json:"tn_w"`
	ThumbnailHeight int   `json:"tn_h"`
	RenamedFilename int64 `json:"tim"`
	ReplyTo         int64 `json:"resto"`
	Sticky          int
	Closed          int
	Now             string
	Tripcode        string `json:"trip"`
	ID              string
	Capcode         string
	Country         string
	CountryName     string `json:"country_name"`
	Email           string
	FileDeleted     int
	Spoiler         int
	CustomSpoiler   int
	OmittedPosts    int
	OmittedImages   int
	BumpLimit       int
	ImageLimit      int
}

func (p Post) Info() string {
	return fmt.Sprintf("Thread #%d: %d replies, %d images\n", p.Number, p.Replies, p.Images)
}

func (p Post) HasAttachment() bool {
	return len(p.Filename) > 0 && len(p.Ext) > 0
}

func (p Post) CompleteFilename() string {
	return fmt.Sprintf("%s%s", p.Filename, p.Ext)
}

func (p Post) UnpackedMD5() (string, error) {
	md5, err := base64.StdEncoding.DecodeString(p.MD5)
	return string(md5), err
}

func (p Post) FullImageUrl() string {
	if !p.HasAttachment() {
		return ""
	}
	return fmt.Sprintf(apiFullImageUrl, p.Board, p.RenamedFilename, p.Ext)
}

func (p Post) ThumbnailUrl() string {
	if !p.HasAttachment() {
		return ""
	}
	return fmt.Sprintf(apiThumbnailImageUrl, p.Board, p.RenamedFilename)
}
