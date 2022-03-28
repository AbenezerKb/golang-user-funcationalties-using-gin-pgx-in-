package db

import (
	"context"
	"fmt"
	"gin-exercise/entity"
	rest_error "gin-exercise/error"
	"os"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	//"google.golang.org/protobuf/internal/errors"
)

var (
	listAll = "SELECT * FROM person;"
	getUser = "SELECT * FROM person WHERE email = $1 AND password = $2"
	insert  = "INSERT INTO person(first_name, second_name, age, email, profile, password, address, phonenumber) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"
)

func connectDB() (c *pgx.Conn, err error) {
	postgres := "postgresql://postgres:abeny@localhost:5432/Exersice"
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

	DbPool, err := connectDB()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer DbPool.Close(context.Background())

	rows, err := DbPool.Query(context.Background(), listAll)
	if err != nil {
		rest_error.NewInternalServerError("error while executing query")
	}

	users := []string{}

	// iterate through the rows
	for rows.Next() {

		values, err := rows.Values()
		if err != nil {
			rest_error.NewInternalServerError("error while iterating dataset from DB")
		}

		users = append(users, fmt.Sprintf("%v", values))
	}

	return users
}

func SaveUser(user entity.User) {

	DbPool, err := connectDB()

	if err != nil {
		rest_error.NewInternalServerError("database error")
		return
	}

	DbPool.Exec(context.Background(), insert, user.FirstName, user.SecondName, user.Age, user.Email, user.Profile, user.Password, user.Address, user.PhoneNumber)

	if err != nil {
		rest_error.NewInternalServerError("error while executing query")
	}
	DbPool.Close(context.Background())
}

func UserInfo(email string, password string) bool {

	DbPool, err := connectDB()

	if err != nil {
		rest_error.NewInternalServerError("database error")
		return false
	}
	defer DbPool.Close(context.Background())

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		rest_error.NewInternalServerError("password confirmation error")
		return false
	}
	rows, err := DbPool.Query(context.Background(), getUser, email, string(hashedPW[:]))
	if err != nil {

		rest_error.NewInternalServerError("database error!")
		return false
	}

	// iterate through the rows
	var user entity.User

	for rows.Next() {

		rows.Scan(&user)
		fmt.Println("user: ", user)
		if err != nil {
			rest_error.NewInternalServerError("database error!")
			return false
		}

	}

	return true
}
