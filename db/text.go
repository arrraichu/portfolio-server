package contentDb

import (
	content "arrraichu/portfolio-server/internal"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func PostTextContent(db *sqlx.DB, input content.Content) error {
	_, err := db.Exec(
		"INSERT INTO content (type, page_path, index, header, textblock1) VALUES ($1, $2, $3, $4, $5)",
		input.Type, input.PagePath, input.Index, input.Header, input.Textblock1,
	)
	if err != nil {
		return fmt.Errorf("PostTextContent: %v", err)
	}

	return nil
}
