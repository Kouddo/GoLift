package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"log"
)

func checkUserNameExists(db *pgx.Conn, userName string) (bool, error) {
	containsQuery := `SELECT EXISTS (
	SELECT 1
	FROM public.users
	WHERE username = $1
);`

	var containsUsername bool

	err := db.QueryRow(context.Background(), containsQuery, userName).Scan(&containsUsername)

	if err != nil {
		log.Fatal(err)
		return false, errors.New("db connection error")
	}

	return containsUsername, nil
}

func addUserToDB(db *pgx.Conn, userName string, PasswordHash []byte) error {

	addQuery := `INSERT INTO public.users(username, password) VALUES ($1, $2)`

	_, err := db.Exec(context.Background(), addQuery, userName, string(PasswordHash))

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
