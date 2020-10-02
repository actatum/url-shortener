package shortener

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	createTable = `
	CREATE TABLE IF NOT EXISTS urls (
		slug varchar(11) PRIMARY KEY,
		url varchar(1000) NOT NULL,
		uses int NOT NULL,
		created_at bigint NOT NULL
	)
	`
)

type repository struct {
	cache map[string]string
	pool  *pgxpool.Pool
}

func newPgxPool() (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(), createTable)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func newRepository() (*repository, error) {
	pool, err := newPgxPool()
	if err != nil {
		return nil, err
	}

	return &repository{
		cache: make(map[string]string),
		pool:  pool,
	}, nil
}

// Create stores a key value pair of slug and url
func (r *repository) Create(req *Request) (*Response, error) {
	_, err := r.pool.Exec(context.Background(), "INSERT INTO urls(slug, url, uses, created_at) VALUES ($1, $2, $3, $4)", req.Slug, req.URL, 0, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	return &Response{Slug: req.Slug, URL: req.URL}, nil
}

// Read reads a shortened url given a key of full url
func (r *repository) Read(req *Request) (*Response, error) {
	var resp Response

	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	_, err = r.pool.Exec(context.Background(), "UPDATE urls SET uses = uses + 1 WHERE slug = $1", req.Slug)
	if err != nil {
		return nil, err
	}

	err = r.pool.QueryRow(context.Background(), "SELECT * FROM urls WHERE slug = $1 ORDER BY slug LIMIT 1", req.Slug).Scan(&resp.Slug, &resp.URL, &resp.Uses, &resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *repository) Top5() ([]Response, error) {
	var resp []Response

	rows, err := r.pool.Query(context.Background(), "SELECT * FROM urls ORDER BY uses DESC LIMIT 5")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r Response
		if err = rows.Scan(&r.Slug, &r.URL, &r.Uses, &r.CreatedAt); err != nil {
			return nil, err
		}

		resp = append(resp, r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
