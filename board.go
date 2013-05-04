package imgchan

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
