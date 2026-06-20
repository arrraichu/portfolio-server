package content

type ContentType struct {
	Code  string `json:"code"`
	Label string `json:"label"`
	Reqs  string `json:"reqs"`
}

type Content struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	PagePath string `json:"page_path"`
	Index    int    `json:"index"`

	Header     string `json:"header"`
	Subheader1 string `json:"subheader1"`
	Subheader2 string `json:"subheader2"`

	Textblock1 string `json:"textblock1"`
	Textblock2 string `json:"textblock2"`

	ButtonText1 string `json:"button1text"`
	ButtonHref1 string `json:"button1href"`
	ButtonText2 string `json:"button2text"`
	ButtonHref2 string `json:"button2href"`
	ButtonText3 string `json:"button3text"`
	ButtonHref3 string `json:"button3href"`
	ButtonText4 string `json:"button4text"`
	ButtonHref4 string `json:"button4href"`

	ImageSrc1 string `json:"image1src"`
	ImageAlt1 string `json:"image1alt"`
	ImageSrc2 string `json:"image2src"`
	ImageAlt2 string `json:"image2alt"`
	ImageSrc3 string `json:"image3src"`
	ImageAlt3 string `json:"image3alt"`
	ImageSrc4 string `json:"image4src"`
	ImageAlt4 string `json:"image4alt"`
}
