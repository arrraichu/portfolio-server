package content

type ContentType struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type Content struct {
	Id         int    `json:"id"`
	PageSource string `json:"page_source"`
}
