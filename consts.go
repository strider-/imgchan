package imgchan

const (
	// Api Endpoints
	apiBoardUrl          = "https://api.4chan.org/boards.json"
	apiBoardPageUrl      = "https://api.4chan.org/%s/%d.json"
	apiCatalogUrl        = "https://api.4chan.org/%s/catalog.json"
	apiThreadUrl         = "https://api.4chan.org/%s/res/%d.json"
	apiFullImageUrl      = "https://images.4chan.org/%s/src/%d%s"
	apiThumbnailImageUrl = "https://thumbs.4chan.org/%s/thumb/%ds.jpg"
	apiSpoilerImageUrl   = "https://static.4chan.org/image/spoiler.png"

	// Regular Expressions
	rxThreadUrl = `/(?P<board>[^/])/res/(?P<thread>\d+)$`

	// Error messages
	errInvalidUrl      = "invalid 4chan url"
	errInvalidThreadId = "invalid 4chan thread id"
	err404             = "specified resource is no longer available"
	errNoImage         = "no attachment for this post"

	// Site Urls
	siteBoardUrl     = "https://boards.4chan.org/%s"
	siteBoardPageUrl = "https://boards.4chan.org/%s/%d"
)
