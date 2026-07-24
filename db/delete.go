package contentDb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DeleteContent(db *sqlx.DB, id string) error {
	result, err := db.Exec("DELETE FROM content WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteContent: %v", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteContent: %v", err)
	}

	if n == 0 {
		return fmt.Errorf("DeleteContent: no rows affected")
	}

	return nil
}
