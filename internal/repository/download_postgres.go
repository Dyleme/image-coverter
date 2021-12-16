package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DownloadPostgres struct {
	db *sql.DB
}

func NewDownloadPostgres(db *sql.DB) *DownloadPostgres {
	return &DownloadPostgres{db: db}
}

func (d *DownloadPostgres) GetImageURL(ctx context.Context, userID, imageID int) (string, error) {
	query := fmt.Sprintf(`SELECT image_url FROM %s WHERE user_id = $1 AND id = $2`, ImageTable)
	row := d.db.QueryRow(query, userID, imageID)

	var urlImage string

	if err := row.Scan(&urlImage); err != nil {
		return "", fmt.Errorf("repo: %w", err)
	}

	return urlImage, nil
}
