package contentDb

import (
	content "arrraichu/portfolio-server/internal"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func FetchAllContent(db *sqlx.DB) ([]content.Content, error) {
	var posts []content.Content

	err := db.Select(&posts, "SELECT * FROM content LIMIT 100")
	if err != nil {
		return nil, err
	}

	return posts, nil
}

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

func UpdateContent(db *sqlx.DB, c content.Content) error {
	if c.Id == 0 {
		return fmt.Errorf("UpdateContent: id is required")
	}

	genericContent, err := toGenericMap(c)
	if err != nil {
		return fmt.Errorf("UpdateContent: %v", err)
	}

	newKeys := make([]string, 0)
	for key, value := range genericContent {
		switch key {
		case "id":
			continue
		default:
			if value != nil {
				newKeys = append(newKeys, fmt.Sprintf("%s = :%s", key, key))
			}
		}
	}

	if len(newKeys) == 0 {
		return fmt.Errorf("UpdateContent: no fields to update")
	}

	query := fmt.Sprintf("UPDATE content SET %s WHERE id=:id", strings.Join(newKeys, ", "))

	result, err := db.NamedExec(query, c)
	if err != nil {
		return fmt.Errorf("UpdateContent: %v", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateContent: %v", err)
	}
	if n == 0 {
		return fmt.Errorf("UpdateContent: no row found with id=%d", c.Id)
	}

	return nil
}

func toGenericMap(c content.Content) (map[string]interface{}, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return m, nil
}
