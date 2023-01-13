package main

import (
	"fmt"
	"library/handler"
	"log"
	"net/http"

	// "database/sql"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var createTable = `
CREATE TABLE IF NOT EXISTS categories (
	id	serial,
	name text,
	status boolean,

	primary key (id)
);

CREATE TABLE IF NOT EXISTS books (
	id	serial,
	category_id integer,
	book_name text,
	author_name text,
	details text,
	image text,
	status boolean,

	primary Key (id)
);

CREATE TABLE IF NOT EXISTS bookings (
	id	serial,
	user_id integer,
	book_id integer,
	start_time timestamp,
	end_time timestamp,

	primary Key (id)
);

CREATE TABLE IF NOT EXISTS users (
	id	serial,
	first_name text,
	last_name text,
	email text,
	password text,
	is_verified boolean,

	primary Key (id)
);`

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=password dbname=library sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(createTable)
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	store := sessions.NewCookieStore([]byte("jsowjpw38eowj4ur82jmaole0uehqpl"))
	r := handler.New(db, decoder, store)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server starting...")
	fmt.Fprint(w, r)
}
