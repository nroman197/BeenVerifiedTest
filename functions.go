package main

import (
	//"fmt"
	"net/http"
	"database/sql"
  "goji.io/pat"
	_ "github.com/mattn/go-sqlite3"
	"encoding/json"
	"strings"
)


// global database
var myDB *sql.DB

// function to get all songs that match with a search title value
func searchSongsbyTitle(w http.ResponseWriter, r *http.Request) {
	// define param of path
	title := pat.Param(r, "title")
	// this is the sql query to get matching songs
	var query ="SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.song LIKE '%"+ title + "%' ORDER BY S.song ASC;"
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
	var query ="SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.artist LIKE '%"+ artist + "%' ORDER BY S.artist ASC;"
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
	var query ="SELECT S.Song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE G.name LIKE '%"+ genre + "%' ORDER BY G.name ASC;"
	// call function to get queried songs from database
	songsList := getSongs(query)
	// write JSON values to an output stream
	json.NewEncoder(w).Encode(songsList)
}

// function to get all songs that match with a search genre value
func searchSongsbyLength(w http.ResponseWriter, r *http.Request) {
	// define param of path
	lenghts := pat.Param(r, "minLength&maxLength")
  minMaxLengths := strings.Split(lenghts, "&")
  minLength := minMaxLengths[0]
  maxLength := minMaxLengths[1]
	// this is the sql query to get matching songs
	var query ="SELECT S.Song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.length BETWEEN "+ minLength + " AND " + maxLength + " ORDER BY S.length ASC;"
	// call function to get queried songs from database
	songsList := getSongs(query)
	// write JSON values to an output stream
	json.NewEncoder(w).Encode(songsList)
}

// function to get numberofsongs and total length grouped by genre
func getListofGenres(w http.ResponseWriter, r *http.Request) {
	// this is the sql query to get matching songs
	var query ="SELECT G.name AS Genre, COUNT(1) AS NumberofSongs, SUM(S.length) AS TotalLength FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre GROUP BY G.name ORDER BY G.name;"
	// call function to get queried songs from database
	songsList := getGenreList(query)
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

// function to get queried list of genres from Database
// params: pQuery string (SQL query)
// returns: []GenreList (array of GenreList struct)
func getGenreList(pQuery string) []GenreList{
	// database connection
	myDB, err := sql.Open("sqlite3", "./jrdd.db")
	// check any database error
	checkErr(err)

	// execute query
	rows, err := myDB.Query(pQuery)
	// check any database error
	checkErr(err)

	// variables to store retrieved data from query
	var genre string
	var numberofSongs int
	var totalLength int
	// array to store all retrieved rows of type GenreList
	var genreList []GenreList

	// for statement to get rows
	for rows.Next() {
		err := rows.Scan(&genre, &numberofSongs, &totalLength)
		checkErr(err)
		// create new Song struct for each row
		new:= GenreList{Genre: genre, NumberofSongs: numberofSongs, TotalLength: totalLength}
		// append new Song struct to songs array
		genreList = append(genreList, new)
	}
	// good habit to close connections
	rows.Close()
	myDB.Close()

	return genreList
}
