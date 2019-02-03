# Been Verified Coding Challenge
### Description

This program is intended to expose searchable data stored in SQLite database via API using json format.

The given database stores a collection of songs in a defined structure, the mission is to make this information searchable by different criteria.

Songs have the following attributes:
* Title
* Artist
* Genre
* Length

Database can be retrieved from:
https://s3.amazonaws.com/bv-challenge/jrdd.db.

#### Functions
* Search songs by title.
* Search songs by artist.
* Search songs by genre.
* Search songs by length given minimum and maximum.
* Search total number of songs and total length by genre.

### Requirements & Installation
This program requires the following:

- [x]  Install Golang 1.9

Recommended commands for ubuntu:
```
sudo add-apt-repository ppa:gophers/archive
sudo apt update
sudo apt-get install golang-1.9-go
```
For setting environment variable, check out this link: https://github.com/golang/go/wiki/SettingGOPATH.

- [x] Goji

Recommended commands for ubuntu:
```
     go get goji.io
```
- [x] Sqllite3

Recommended commands for ubuntu:
```
go get github.com/mattn/go-sqlite3
```
### How to use

To  run this program open Terminal and run ./BeenVerifiedTest.

Go to browser and search any of the following functions:

#### Search song by title
This function will return every found song whose title matches with the search value written in the path.

Example:
http://localhost:8000/searchSongsbyTitle/Hey%20Jude

#### Search song by artist
This function will return every found song whose artist matches with the search value written in the path.

Example:
http://localhost:8000/searchSongsbyArtist/The%20Black%20Eyed

#### Search song by genre
This function will return every found song whose genre matches with the search value written in the path.

Example:
http://localhost:8000/searchSongsbyGenre/Classic

#### Search song by length
This function will return every found song whose length is in the range written in the path. Format is: minLenght&maxLength.

Example:
http://localhost:8000/searchSongsbyLength/100&200

#### Get List of genres
This function will return number of songs and total length by each genre.

Example:

http://localhost:8000/getListofGenres/
