package repository

import (
	"context"
	"fmt"
	"os"
	"reflect"

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

// Usage: tagMap(&model.User{})
func tagMap(v interface{}) map[string]string {
	m := make(map[string]string)

	r := reflect.TypeOf(v).Elem()
	for i := 0; i < r.NumField(); i++ {
		json := r.Field(i).Tag.Get("json")
		db := r.Field(i).Tag.Get("db")
		m[json] = db
	}

	return m
}
