package main //import "github.com/felipejfc/go-pgbouncer-rr-example"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-pg/pg"
)

var db *pg.DB

func main() {

	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5434",
		User:     "postgres",
		Password: "password",
		Database: "example",
	})
	defer db.Close()

	err := createSchema()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected to db successfully!")
	}

	router := NewRouter()

	fmt.Println("will start listening to :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
