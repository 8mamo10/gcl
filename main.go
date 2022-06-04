package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(mysql:3306)/mrhc")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(w, `{"error": "failed to query users: %v"}`, err)
		return
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var id int
		var account_id int
		var name string
		err := rows.Scan(&id, &account_id, &name)
		if err != nil {
			fmt.Fprintf(w, `{"error": "failed to retrieve data: %v"}`, err)
			return
		}
		users = append(users, fmt.Sprintf(`{"id": %d, "account_id: %d, "name": "%s"}`, id, account_id, name))
	}
	fmt.Fprintf(w, "["+strings.Join(users, ",")+"]")
}
