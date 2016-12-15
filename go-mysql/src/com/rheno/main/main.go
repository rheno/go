package main

import (

"fmt"
_"github.com/go-sql-driver/mysql"
"database/sql"

)

func main() {

	// My MAMP Mysql is on port 8889 and there is mysqlgo database 
	// Don't forget to add user table with id, username and password as columns
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/mysqlgo?autocommit=true")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Check database connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// SELECT All Rows
	rows, err := db.Query("SELECT username, password FROM user")

	if err != nil {
		panic(err.Error())
	}


	for rows.Next() {
		var username string
		var password string
		err = rows.Scan(&username, &password)

		fmt.Println(username," ",password)

		if err != nil {
			panic(err.Error())
		}
	}

	// Select One or More Row
	var name string

	err = db.QueryRow("SELECT username FROM user WHERE id = ?", 1).Scan(&name)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(name)


	// Insert 
	_, err = db.Exec("INSERT INTO user(username, password) VALUES(?, ?)", "rheno", "ganteng")

        if err != nil {
		panic(err.Error())
        }

	fmt.Println("insert done !!!")

	// Update
	_, err = db.Exec("UPDATE user SET username = ? WHERE password = ?", "keren", "ganteng")


        if err != nil {
		panic(err.Error())
        }

	fmt.Println("update done !!!")


	// Delete
	_, err = db.Exec("DELETE FROM user WHERE username = ?", "keren")


        if err != nil {
		panic(err.Error())
        }

	fmt.Println("delete done !!!")


}
