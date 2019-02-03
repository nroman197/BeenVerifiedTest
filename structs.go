package main

import (
	"net/http"
	"database/sql"
)

// global database
var myDB *sql.DB

// definition of Song struct for retrieved data
type Song struct {
	Title string "json:title, omitempty"
	Artist string "json:artist, omitempty"
	Genre string "json:string, omitempty"
	Length int "json:length, omitempty"
}

// definition of GenreList struct for retrieved data
type GenreList struct {
	Genre string "json:genre, omitempty"
	NumberofSongs int "json:numberofsongs, omitempty"
	TotalLength int "json:totallength, omitempty"
}

// this struct is used to define each different function and its path to be called in the main function
type Path struct {
	Pat string
	Function http.HandlerFunc
}

// array of Paths to define all required get functions
type PathList []Path

// in case it is needed to add one or more functions, just add it to the structure below
var plist = PathList{
		Path{
			"/searchSongsbyTitle/:title",
			searchSongsbyTitle,
		},
		Path{
			"/searchSongsbyArtist/:artist",
			searchSongsbyArtist,
		},
		Path{
			"/searchSongsbyGenre/:genre",
			searchSongsbyGenre,
		},
		Path{
			"/searchSongsbyLength/:minLength&maxLength",
			searchSongsbyLength,
		},
		Path{
			"/getListofGenres/",
			getListofGenres,
		},
}
