package imgchan

import "fmt"

type Post struct {
	Number          int `json:"no"`
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
