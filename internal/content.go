package content

type ContentType struct {
	Code  string `json:"code"`
	Label string `json:"label"`
	Reqs  string `json:"reqs"`
}

type Content struct {
	Id         int    `json:"id"`
	PageSource string `json:"page_source"`
}
