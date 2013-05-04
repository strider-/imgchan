package imgchan

import "fmt"

type Page struct {
	Index   int `json:"page"`
	Threads []Post
}

func (p Page) Info() string {
	return fmt.Sprintf("Page #%d, %d threads\n", p.Index, len(p.Threads))
}
