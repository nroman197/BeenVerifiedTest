package main

// definition of song struct for retrieved data
type Song struct {
	Title string "json:title"
	Artist string "json:artist"
	Genre string "json:string"
	Length int "json:length"
}
