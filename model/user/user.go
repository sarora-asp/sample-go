package userModel

import (
	"fmt"
	"log"
	userpb "sample/twirp/rpc/user"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AuthUser struct {
	Id   int    `json:"user_id"`
	Name string `json:"name"`
}

func InsertOne(db sqlx.DB, user *userpb.User) int64 {
	fmt.Println("USER", user)
	n, err := db.MustExec(`INSERT INTO users (name, email, password) VALUES ($1, $2, $3);`, user.Name, user.Email, user.Password).RowsAffected()
	if err != nil {
		log.Fatalln("Unable to insert shit", err)
	}
	return n
}

func FindUserByEmail(db sqlx.DB, email string) *User {
	rows, err := db.Queryx(`SELECT * FROM users WHERE email = $1 `, email)
	if err != nil {
		fmt.Println(err)
	}
	var user User
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln("Something went wrong", err)
		}
	}

	return &user
}

func FindUserById(db sqlx.DB, id int) *User {
	rows, err := db.Queryx(`SELECT * FROM users WHERE email = $1 `, id)
	if err != nil {
		log.Fatalln("Unable to insert shit", err)
	}
	var user User
	for rows.Next() {
		fmt.Println("NROWS", rows)
		err = rows.StructScan(&user)
		fmt.Println("ERROR IN SCAN", err)
	}
	return &user
}
