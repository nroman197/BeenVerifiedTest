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

// definition of song struct for retrieved data
type Song struct {
	Artist string "json:artist"
	Song string "json:song"
	Genre string "json:string"
	Length int "json:length"
}

// function to get all songs that match with a search value
func searchSong(w http.ResponseWriter, r *http.Request) {
	search := pat.Param(r, "search")
	// this is the sql query to get matching songs
	var query ="SELECT S.artist AS Artist, S.Song AS Song, G.name AS 'Genre', S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.Song LIKE '%"+ search + "%';"
	// database connection
	myDB, err := sql.Open("sqlite3", "./jrdd.db")
	// check any database error
	checkErr(err)

	// execute query
	rows, err := myDB.Query(query)
	// check any database error
	checkErr(err)

	// variables to store retrieved data from query
	var artist string
	var song string
	var genre string
	var length int
	// array to store all retrieved rows of type Song
	var songsList []Song

	// for statement to get rows
	for rows.Next() {
		err := rows.Scan(&artist, &song, &genre, &length)
		checkErr(err)
		// create new Song struct for each row
		new:= Song{Artist: artist, Song: song, Genre: genre, Length: length}
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
	mux.HandleFunc(pat.Get("/searchSong/:search"), searchSong)

	// listening port
	http.ListenAndServe("localhost:8000", mux)
}

// function to check errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
