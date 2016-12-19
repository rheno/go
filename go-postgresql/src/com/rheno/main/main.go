package main


import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:1q2w3e4r5t6y@localhost/iseng?sslmode=disable")

	if err != nil {
		panic(err.Error())
	}


	/* INSERT */
	_, err = db.Exec(`INSERT INTO users(username, password) VALUES('rheno', 'cool')`)


	if err != nil {
		panic(err.Error())
	}



	/* UPDATE */
	_, err = db.Exec(`UPDATE users SET username = 'nice', password = 'strong' WHERE id = 1`)


	if err != nil {
		panic(err.Error())
	}



	/* DELETE */
	_, err = db.Exec(`DELETE FROM users WHERE id = 1`)

	if err != nil {
		panic(err.Error())
	}



	/* SELECT */
	rows, err := db.Query("SELECT username, password FROM users")

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


}
