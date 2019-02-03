package main

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
