package main

import "testing"

func TestArtistQueries(t *testing.T) {
	tables := []struct {
    value string
    query string
		artist string
	}{
		{"Beatl", "SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.artist LIKE '%Beatl%' ORDER BY S.artist ASC;", "Beatles"},
		{"The Black", "SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.artist LIKE '%The Black%' ORDER BY S.artist ASC;", "The Black Eyed Peas"},
	}

	for _, table := range tables {
		result := getSongs(table.query)
		if result[0].Artist != table.artist {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result[0].Artist, table.artist)
		}
	}
}
