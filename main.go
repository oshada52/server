package main

import (
	"fmt"
	"net/http"
	"os"
	db "server/db"
)

func getNames(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Connection.Query("SELECT name FROM locations")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	data := ""

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name)
		data += fmt.Sprintf("%s ", name)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, data)
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {

	db.InitDB()
	defer db.Connection.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/data", getData)
	http.HandleFunc("/names", getNames)

	serverEnv := os.Getenv("SERVER_ENV")

	if serverEnv == "DEV" {
		http.ListenAndServe(":8080", nil)
	} else if serverEnv == "PROD" {
		http.ListenAndServeTLS(
			":443",
			"/etc/letsencrypt/live/ceyloninfo.site/fullchain.pem",
			"/etc/letsencrypt/live/ceyloninfo.site/privkey.pem",
			nil,
		)
	}

}
