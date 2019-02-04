package main

import "testing"

func TestGenreQueries(t *testing.T) {
	tables := []struct {
    value string
    query string
		genre string
	}{
		{"Classic", "SELECT S.Song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE G.name LIKE '%Classic%' ORDER BY G.name ASC;", "Classic Rock"},
		{"Rap", "SELECT S.Song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE G.name LIKE '%Rap%' ORDER BY G.name ASC;", "Rap"},

	}

	for _, table := range tables {
		result := getSongs(table.query)
		if result[0].Genre != table.genre {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result[0].Genre, table.genre)
		}
	}
}
