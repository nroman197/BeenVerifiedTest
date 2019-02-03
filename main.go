package main

import (
	//"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
)

func main() {
	// create mux
	mux := goji.NewMux()
	// define path and function to be called
	mux.HandleFunc(pat.Get("/searchSongsbyTitle/:title"), searchSongsbyTitle)
	mux.HandleFunc(pat.Get("/searchSongsbyArtist/:artist"), searchSongsbyArtist)
	mux.HandleFunc(pat.Get("/searchSongsbyGenre/:genre"), searchSongsbyGenre)
	// listening port
	http.ListenAndServe("localhost:8000", mux)
}