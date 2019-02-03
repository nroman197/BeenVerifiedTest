package main

import (
	//"fmt"
	"net/http"
	"database/sql"
  "goji.io/pat"
	_ "github.com/mattn/go-sqlite3"
	"encoding/json"
	//"strings"
)


// global database
var myDB *sql.DB

// function to get all songs that match with a search title value
func searchSongsbyTitle(w http.ResponseWriter, r *http.Request) {
	// define param of path
	title := pat.Param(r, "title")
	// this is the sql query to get matching songs
	var query ="SELECT S.Song AS Title,  S.artist AS Artist, G.name AS 'Genre', S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.Song LIKE '%"+ title + "%';"
	// call function to get queried songs from database
	songsList := getSongs(query)
	// write JSON values to an output stream
	json.NewEncoder(w).Encode(songsList)
}

// function to get all songs that match with a search artist value
func searchSongsbyArtist(w http.ResponseWriter, r *http.Request) {
	// define param of path
	artist := pat.Param(r, "artist")
	// this is the sql query to get matching songs
	var query ="SELECT S.Song AS Title,  S.artist AS Artist, G.name AS 'Genre', S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.artist LIKE '%"+ artist + "%';"
	// call function to get queried songs from database
	songsList := getSongs(query)
	// write JSON values to an output stream
	json.NewEncoder(w).Encode(songsList)
}

// function to get all songs that match with a search genre value
func searchSongsbyGenre(w http.ResponseWriter, r *http.Request) {
	// define param of path
	genre := pat.Param(r, "genre")
	// this is the sql query to get matching songs
	var query ="SELECT S.Song AS Title,  S.artist AS Artist, G.name AS 'Genre', S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE G.name LIKE '%"+ genre + "%';"
	// call function to get queried songs from database
	songsList := getSongs(query)
	// write JSON values to an output stream
	json.NewEncoder(w).Encode(songsList)
}

// function to get queried songs from Database
// params: pQuery string (SQL query)
// returns: []Song (array of Song struct)
func getSongs(pQuery string) []Song{
	// database connection
	myDB, err := sql.Open("sqlite3", "./jrdd.db")
	// check any database error
	checkErr(err)

	// execute query
	rows, err := myDB.Query(pQuery)
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

	return songsList
}
