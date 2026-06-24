package contentDb

import (
	content "arrraichu/portfolio-server/internal"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func PostTitle(db *sqlx.DB, input content.Content) error {
	_, err := db.Exec(
		"INSERT INTO content (type, page_path, index, header, subheader1, textblock1) VALUES ($1, $2, $3, $4, $5, $6)",
		input.Type, input.PagePath, input.Index,
		input.Header, input.Subheader1, input.Textblock1,
	)
	if err != nil {
		return fmt.Errorf("PostTitle: %v\n", err)
	}

	return nil
}

func PostTextContent(db *sqlx.DB, input content.Content) error {
	_, err := db.Exec(
		"INSERT INTO content (type, page_path, index, header, textblock1) VALUES ($1, $2, $3, $4, $5)",
		input.Type, input.PagePath, input.Index,
		input.Header, input.Textblock1,
	)
	if err != nil {
		return fmt.Errorf("PostTextContent: %v\n", err)
	}

	return nil
}

func PostTextButtonsContent(db *sqlx.DB, input content.Content) error {
	_, err := db.Exec(
		"INSERT INTO content (type, page_path, index, header, textblock1, button1text, button1href, button2text, button2href) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		input.Type, input.PagePath, input.Index,
		input.Header, input.Textblock1,
		input.ButtonText1, input.ButtonHref1,
		input.ButtonText2, input.ButtonHref2,
	)
	if err != nil {
		return fmt.Errorf("PostTextButtonsContent: %v\n", err)
	}

	return nil
}
