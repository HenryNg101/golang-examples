package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func construct_tables(ctx context.Context, conn *pgx.Conn) {
	_, err := conn.Exec(ctx, `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		createdAt TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(ctx, `
	CREATE TABLE orders (
		id SERIAL PRIMARY KEY,
		userId INT REFERENCES users(id),
		price INT,
		createdAt TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(ctx, `
	CREATE TABLE order_items (
		id SERIAL PRIMARY KEY,
		orderId INT REFERENCES orders(id),
		price INT,
		name VARCHAR(30),
		createdAt TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}
}

func generate_records(ctx context.Context, conn *pgx.Conn) {
	_, err := conn.Exec(ctx, `
	INSERT INTO users(name, createdAt)
	SELECT
		'user_' || g,
		NOW() - (g || ' seconds')::interval
	FROM generate_series(1,1000000) g
	`)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(ctx, `
	INSERT INTO orders(userId, price, createdAt)
	SELECT
		(floor(random()*1000000))::int + 1,
		(floor(random()*500))::int,
		NOW() - (g || ' seconds')::interval
	FROM generate_series(1,5000000) g
	`)
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec(ctx, `
	INSERT INTO order_items(orderId, price, name, createdAt)
	SELECT
		(floor(random()*5000000))::int + 1,
		(floor(random()*500))::int,
		'default',
		NOW() - (g || ' seconds')::interval
	FROM generate_series(1,10000000) g
	`)
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := pgx.Connect(context.Background(),
		"postgresql://postgres:secret@localhost:5432")
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	// tag, err := conn.Exec(context.Background(), "CREATE DATABASE mydb")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%T %s\n", tag, tag)

	construct_tables(context.Background(), conn)
	generate_records(context.Background(), conn)

	// _, err = conn.Exec(context.Background(), "DROP DATABASE mydb")
	// if err != nil {
	// 	panic(err)
	// }
}
