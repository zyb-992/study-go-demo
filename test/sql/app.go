package main

import (
	"database/sql"
)

func AddRecord(db *sql.DB, userId, ProductId int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	if _, err = tx.Exec("UPDATE products SET views = views+1"); err != nil {
		return nil
	}

	if _, err = tx.Exec(
		"INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)",
		userId, ProductId); err != nil {
		return err
	}
	return nil
}
