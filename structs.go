package main

// definition of song struct for retrieved data
type Song struct {
	Title string "json:title, omitempty"
	Artist string "json:artist, omitempty"
	Genre string "json:string, omitempty"
	Length int "json:length, omitempty"
}
