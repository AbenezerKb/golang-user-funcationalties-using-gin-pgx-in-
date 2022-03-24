package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func connectDB() (c *pgx.Conn, err error) {
	postgres := "postgresql://postgres:abeny@localhost:5432/goapi"
	conn, err := pgx.Connect(context.Background(), postgres)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database:\n\t%v\n", err.Error())
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping:\n\t%v\n", err.Error())
		return nil, err
	}
	return conn, err
}

func Userslist() []string {

	databaseUrl := "postgresql://postgres:abeny@localhost:5432/Exersice"

	// this returns connection pool
	DbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer DbPool.Close()

	rows, err := DbPool.Query(context.Background(), "select * from person")
	if err != nil {
		log.Fatal("error while executing query")
	}

	users := []string{}

	// iterate through the rows
	for rows.Next() {
		//var r entity.User
		// err := rows.Scan(&r.FirstName, &r.SecondName, &r.Age, &r.Email)
		// fmt.Println(r)
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		users = append(users, fmt.Sprintf("%v", values))
	}

	return users
}

func users(email string, password string) bool {

	databaseUrl := "postgresql://postgres:abeny@localhost:5432/Exersice"

	// this returns connection pool
	DbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer DbPool.Close()

	rows, err := DbPool.Query(context.Background(), "select email, password from person where email = %v AND password = %v", email, password)
	if err != nil {
		log.Fatal("error while executing query")
	}

	users := []string{}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
			return false
		}
		fmt.Println("values: ", values)

		users = append(users, fmt.Sprintf("%v", values))
	}

	return true
}