package contentDb

import (
	content "arrraichu/portfolio-server/internal"

	"github.com/jmoiron/sqlx"
)

func FetchContent(db *sqlx.DB, path string) ([]content.Content, error) {
	var posts []content.Content

	err := db.Select(&posts, "SELECT * FROM content WHERE page_path = $1", path)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func FetchContentTypes(db *sqlx.DB) ([]content.ContentType, error) {
	var types []content.ContentType

	err := db.Select(&types, "SELECT * FROM content_type")
	if err != nil {
		return nil, err
	}

	return types, nil
}
