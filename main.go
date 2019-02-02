package main

import (
	//"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"encoding/json"
	//"strings"
)

// database
var myDB *sql.DB

// function to get all songs that match with a search value
func searchSongbyTitle(w http.ResponseWriter, r *http.Request) {
	// define param of path
	search := pat.Param(r, "search")
	// this is the sql query to get matching songs
	var query ="SELECT S.Song AS Title,  S.artist AS Artist, G.name AS 'Genre', S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.Song LIKE '%"+ search + "%';"
	// database connection
	myDB, err := sql.Open("sqlite3", "./jrdd.db")
	// check any database error
	checkErr(err)

	// execute query
	rows, err := myDB.Query(query)
	// check any database error
	checkErr(err)

	// variables to store retrieved data from query
	var title string
	var artist string
	var genre string
	var length int
	// array to store all retrieved rows of type Song
	var songsList []Song

	// for statement to get rows
	for rows.Next() {
		err := rows.Scan(&title, &artist, &genre, &length)
		checkErr(err)
		// create new Song struct for each row
		new:= Song{Title: title, Artist: artist, Genre: genre, Length: length}
		// append new Song struct to songs array
		songsList = append(songsList, new)
	}
	// good habit to close connections
	rows.Close()
	myDB.Close()

	// write JSON values to an output stream
	json.NewEncoder(w).Encode(songsList)

}

func main() {
	// create mux
	mux := goji.NewMux()
	// define path and function to be called
	mux.HandleFunc(pat.Get("/searchSongbyTitle/:search"), searchSongbyTitle)

	// listening port
	http.ListenAndServe("localhost:8000", mux)
}
