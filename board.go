package imgchan

import "fmt"

type Board struct {
	Board          string
	Title          string
	Worksafe       int `json:"ws_board"`
	ThreadsPerPage int `json:"per_page"`
	Pages          int
}

func (b Board) IsWorksafe() bool {
	return b.Worksafe == 1
}

func (b Board) Url() string {
	return fmt.Sprintf(siteBoardUrl, b.Board)
}

func (b Board) PageUrl(page int) string {
	return fmt.Sprintf(siteBoardPageUrl, b.Board, page)
}
