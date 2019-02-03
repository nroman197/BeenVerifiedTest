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

* Install Golang 1.9

Recommended commands for ubuntu:
```
sudo add-apt-repository ppa:gophers/archive
sudo apt update
sudo apt-get install golang-1.9-go
export PATH=$PATH:/usr/local/go/bin
```

* Goji

Recommended commands for ubuntu:
```
     go get goji.io
```
* Sqllite3

Recommended commands for ubuntu:
```
go get github.com/mattn/go-sqlite3
```
### How to use
