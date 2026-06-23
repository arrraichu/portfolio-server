package content

type ContentType struct {
	Code  string `json:"code"   db:"code"`
	Label string `json:"label"  db:"label"`
	Reqs  string `json:"reqs"   db:"reqs"`
}

type Content struct {
	Id       int    `json:"id"         db:"id"`
	Type     string `json:"type"       db:"type"`
	PagePath string `json:"page_path"  db:"page_path"`
	Index    int    `json:"index"      db:"index"`

	Header     *string `json:"header,omitempty"      db:"header"`
	Subheader1 *string `json:"subheader1,omitempty"  db:"subheader1"`
	Subheader2 *string `json:"subheader2,omitempty"  db:"subheader2"`

	Textblock1 *string `json:"textblock1,omitempty"  db:"textblock1"`
	Textblock2 *string `json:"textblock2,omitempty"  db:"textblock2"`

	ButtonText1 *string `json:"button1text,omitempty"  db:"button1text"`
	ButtonHref1 *string `json:"button1href,omitempty"  db:"button1href"`
	ButtonText2 *string `json:"button2text,omitempty"  db:"button2text"`
	ButtonHref2 *string `json:"button2href,omitempty"  db:"button2href"`
	ButtonText3 *string `json:"button3text,omitempty"  db:"button3text"`
	ButtonHref3 *string `json:"button3href,omitempty"  db:"button3href"`
	ButtonText4 *string `json:"button4text,omitempty"  db:"button4text"`
	ButtonHref4 *string `json:"button4href,omitempty"  db:"button4href"`

	ImageSrc1 *string `json:"image1src,omitempty"  db:"image1src"`
	ImageAlt1 *string `json:"image1alt,omitempty"  db:"image1alt"`
	ImageSrc2 *string `json:"image2src,omitempty"  db:"image2src"`
	ImageAlt2 *string `json:"image2alt,omitempty"  db:"image2alt"`
	ImageSrc3 *string `json:"image3src,omitempty"  db:"image3src"`
	ImageAlt3 *string `json:"image3alt,omitempty"  db:"image3alt"`
	ImageSrc4 *string `json:"image4src,omitempty"  db:"image4src"`
	ImageAlt4 *string `json:"image4alt,omitempty"  db:"image4alt"`
}
