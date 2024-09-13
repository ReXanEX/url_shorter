package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Функции для работы с базой данных

func AddSortLink(shortLink string, originalLink string) {
	connStr := getConnString()
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close(context.Background())

	var query = fmt.Sprintf("INSERT INTO urls (short, original) VALUES ('%s', '%s');", shortLink, originalLink)

	_, err = db.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("Error inserting new urls")
	}
	defer db.Close(context.Background())
}

func GetOriginalLink(shortLink string) (string, error) {
	connStr := getConnString()
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return "", err
	}
	defer db.Close(context.Background())

	var originalLink string
	err = db.QueryRow(context.Background(), "SELECT original FROM urls WHERE short = $1", shortLink).Scan(&originalLink)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("no record found for shortLink: %s", shortLink)
		}
		return "", err
	}

	return originalLink, nil
}
