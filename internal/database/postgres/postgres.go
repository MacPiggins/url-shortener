package postgres

import (
	"context"
	"url-shortener/internal/database"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

func (db *PostgresDB) Set(token, link string) error {
	_, err := db.pool.Exec(
		context.Background(),
		"INSERT INTO links (token, link) VALUES ($1, $2)",
		token,
		link)

	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			if pgerr.Code == pgerrcode.UniqueViolation {
				return database.UniqueError{}
			}
		}
		return err
	}
	return nil
}

func (db *PostgresDB) Get(token string) (string, error) {
	var link string
	row := db.pool.QueryRow(context.Background(), "SELECT link FROM links WHERE token = $1", token)
	err := row.Scan(&link)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", database.NotFoundError{}
		}
		return "", err
	}
	return link, err
}

func New(link string) (*PostgresDB, error) {
	db := PostgresDB{}
	var err error
	db.pool, err = pgxpool.New(context.Background(), link)
	return &db, err
}

func (db *PostgresDB) Close() {
	db.pool.Close()
}
