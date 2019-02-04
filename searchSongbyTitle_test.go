package main

import "testing"

func TestTitleQueries(t *testing.T) {
	tables := []struct {
    value string
    query string
		title string
	}{
		{"Jude", "SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.song LIKE '%Jude%' ORDER BY S.song ASC;", "Hey Jude"},
		{"Feel", "SELECT S.song AS Title,  S.artist AS Artist, G.name AS Genre, S.length AS Length FROM Songs AS S INNER JOIN Genres AS G ON G.ID = S.genre WHERE S.song LIKE '%Feel%' ORDER BY S.song ASC;", "I Gotta Feeling"},
	}

	for _, table := range tables {
		result := getSongs(table.query)
		if result[0].Title != table.title {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result[0].Title, table.title)
		}
	}
}
