package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func Create() (*Repository, error) {

	pool, err := pgxpool.Connect(context.TODO(), os.Getenv("DB_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil, err
	}

	fmt.Println("Connected to database")

	return &Repository{pool: pool}, err
}
