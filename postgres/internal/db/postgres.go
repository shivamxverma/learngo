package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	connString := "postgresql://shivamverma@localhost:5432/todo_app"

	pool, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		fmt.Println("Error in Conncting to DB")
		return nil, err
	}

	err = pool.Ping(context.Background())

	if err != nil {
		fmt.Println("Error on connecting to db")
		return nil, err
	}

	return pool, nil
}


