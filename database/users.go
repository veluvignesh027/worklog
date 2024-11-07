package database

import (
	"database/sql"
	"fmt"
	"log"
)

func PutIntoUserTable(db *sql.DB, dbname string, user User) error {
	first_name := user.FName
	last_name := user.LName
	email := user.Email
	role := user.Role
	customer_id := user.CustomerId
	created_time := user.CreatedAt
	modified_time := user.ModifiedAt

	_, err := db.Exec("INSERT INTO  users_table(first_name,last_name,email,password,role,customer_id,created_time,modified_time) VALUES($1,$2,$3,$4,$5,$6,$7,$8)",
		first_name,
		last_name,
		email,
		role,
		customer_id,
		created_time,
		modified_time,
	)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		log.Println("User Added into the DB!", first_name)
	}

	return err

}

func PutIntoSessionTable(db *sql.DB, dbname string, user User) error {
	first_name := user.FName
	last_name := user.LName
	email := user.Email
	role := user.Role
	customer_id := user.CustomerId
	created_time := user.CreatedAt
	modified_time := user.ModifiedAt

	_, err := db.Exec("INSERT INTO  users_table(first_name,last_name,email,password,role,customer_id,created_time,modified_time) VALUES($1,$2,$3,$4,$5,$6,$7,$8)",
		first_name,
		last_name,
		email,
		role,
		customer_id,
		created_time,
		modified_time,
	)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		log.Println("User Added into the DB!", first_name)
	}

	return err

}

func PutIntoUserStoryTable(db *sql.DB, dbname string, user User) error {
	first_name := user.FName
	last_name := user.LName
	email := user.Email
	role := user.Role
	customer_id := user.CustomerId
	created_time := user.CreatedAt
	modified_time := user.ModifiedAt

	_, err := db.Exec("INSERT INTO  users_table(first_name,last_name,email,password,role,customer_id,created_time,modified_time) VALUES($1,$2,$3,$4,$5,$6,$7,$8)",
		first_name,
		last_name,
		email,
		role,
		customer_id,
		created_time,
		modified_time,
	)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		log.Println("User Added into the DB!", first_name)
	}

	return err

}
