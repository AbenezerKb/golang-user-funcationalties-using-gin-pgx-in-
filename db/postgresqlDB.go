package db

import (
	"context"
	"fmt"
	"gin-exercise/entity"
	rest_error "gin-exercise/error"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	//"google.golang.org/protobuf/internal/errors"
)

var (
	listAll = "SELECT * FROM person;"
	getUser = "SELECT * FROM person WHERE email = $1 AND password = $2"
	insert  = "INSERT INTO person(first_name, second_name, age, email, profile, password, address, phonenumber) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"
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
		rest_error.NewInternalServerError("error while executing query")
	}

	users := []string{}

	// iterate through the rows
	for rows.Next() {
		//var r entity.User
		// err := rows.Scan(&r.FirstName, &r.SecondName, &r.Age, &r.Email)
		// fmt.Println(r)
		values, err := rows.Values()
		if err != nil {
			rest_error.NewInternalServerError("error while iterating dataset from DB")
		}

		users = append(users, fmt.Sprintf("%v", values))
	}

	return users
}

func SaveUser(user entity.User) {

	databaseUrl := "postgresql://postgres:abeny@localhost:5432/Exersice"

	// this returns connection pool
	DbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer DbPool.Close()
	fmt.Println(user)
	result, err := DbPool.Exec(context.Background(), insert, "INSERT INTO person(first_name, second_name, age, email, profile, password) VALUES($1,$2,$3,$4,$5,$6)")
	//user.FirstName, user.SecondName, user.Age, user.Email, user.Profile, user.Password, user.Address, user.PhoneNumber
	//
	log.Print(result)
	if err != nil {
		rest_error.NewInternalServerError("error while executing query")
	}

	// users := []string{}

	// iterate through the rows
	// for rows.Next() {
	//var r entity.User
	// err := rows.Scan(&r.FirstName, &r.SecondName, &r.Age, &r.Email)
	// fmt.Println(r)
	// values, err := rows.Values()
	// if err != nil {
	// 	log.Fatal("error while iterating dataset")
	// }

	// users = append(users, fmt.Sprintf("%v", values))
	// }

	//	return users
}

func UserInfo(email string, password string) bool {

	databaseUrl := "postgresql://postgres:abeny@localhost:5432/Exersice"

	// this returns connection pool
	DbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {

		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		rest_error.NewInternalServerError("database error!")
		return false
	}

	defer DbPool.Close()
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	rows, err := DbPool.Query(context.Background(), "select * from person where email = $1 AND password = $2 ", email, string(hashedPW[:]))
	if err != nil {

		rest_error.NewInternalServerError("database error!")
		return false
	}

	//users := []string{}

	// iterate through the rows
	var user entity.User
	fmt.Println("after quesry")
	for rows.Next() {
		//values, err := rows.Values()
		rows.Scan(&user)
		fmt.Println("user: ", user)
		if err != nil {
			rest_error.NewInternalServerError("database error!")
			return false
		}
		//fmt.Println("values: ", values)

		//users = append(users, fmt.Sprintf("%v", values))
	}

	return true
}
