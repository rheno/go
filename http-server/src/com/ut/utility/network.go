package utility


import (

"net/http"
"encoding/json"
"database/sql"
_ "github.com/go-sql-driver/mysql"
"strconv"

)


type Response struct {

	Code		int		`json:"code"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data"`

}

type User struct {
	Id		int64		`json:"id"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
}

func GetAllUser(query string) []User {


	var u []User
	db, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(localhost:3306)/mydatabase")

        if err != nil {
                return nil
        }

        defer db.Close()

        rows, err := db.Query(query)

        if err != nil {
                return nil
        }

        defer rows.Close()

        var us User

        for rows.Next() {


                err := rows.Scan(&us.Id, &us.Username, &us.Password)

                if err != nil {
                        return nil
                }
                u = append(u, us)

        }

        err = rows.Err()
        if err != nil {
                return nil
        }



        return u


}

func GetUser(query string, args ...interface{}) *User {

	var usr *User

        db, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(localhost:3306)/mydatabase")

        if err != nil {
                return usr
        }

        defer db.Close()

	var u *User = new(User)

        err = db.QueryRow(query, args[0]).Scan(&(*u).Id, &(*u).Username, &(*u).Password)


        if err != nil {

		if u == nil {
			return nil
		}
                return usr
        }

	usr = u


        return usr

}


func InsertUser(query string, arg ...interface{}) bool {

	db, err := sql.Open("mysql", "USERNAME:PASSWROD@tcp(localhost:3306)/mydatabase")

	if err != nil {
		return false
	}
	defer db.Close()

	ins, err := db.Prepare(query)

	if err != nil {
		return false
	}
	defer ins.Close()

	_, err = ins.Exec(arg[0], arg[1])

	if err != nil {
		return false
	}

	return true

}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {

		w.WriteHeader(404)
		json.NewEncoder(w).Encode(Response{404, "Not Found", nil})
		return

	}

	ch := make(chan bool)

	var data bool

	go func() {

		ch := ch

		ch <- InsertUser("INSERT INTO user(username, password) VALUES( ?, ? )", r.FormValue("username"), r.FormValue("password"))

	}()

	data =<- ch

	w.WriteHeader(200)

	if !data {
		json.NewEncoder(w).Encode(Response{9000, "failed", false})
		return
	}


	json.NewEncoder(w).Encode(Response{0, "Success", true})


}

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {

		w.WriteHeader(404)
		json.NewEncoder(w).Encode(Response{404, "Not Found", nil})
		return

	}

	ch := make(chan []User)

	var data []User

	go func() {

		ch := ch

		ch <- GetAllUser("SELECT * FROM user")

	}()

	data =<- ch

	w.WriteHeader(200)

	if data == nil {
		json.NewEncoder(w).Encode(Response{9001, "data not found", nil})
		return
	}


	json.NewEncoder(w).Encode(Response{0, "Success", data})


}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json")

        if r.Method != "GET" {
                w.WriteHeader(404)
                json.NewEncoder(w).Encode(Response{404, "Not Found", nil})
                return
        }

        w.WriteHeader(200)

        ch := make(chan *User)

        id, _ := strconv.Atoi(r.URL.Query().Get("id"))

        go func() {

		ch := ch

                ch <- GetUser("SELECT * FROM user WHERE id = ?", id)

        }()

        data := <-ch

        if data == nil {
                json.NewEncoder(w).Encode(Response{200, "Data not found", nil})
                return
        }

        json.NewEncoder(w).Encode(Response{200, "Success", data})

}


func CreateServer() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(Response{404, "Not Found", nil})

	})

	mux.HandleFunc("/user/add", AddUserHandler)
	mux.HandleFunc("/user/all", GetAllUserHandler)
	mux.HandleFunc("/user", GetUserHandler)

	http.ListenAndServe(":8080", mux)

}
