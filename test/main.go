package main

import (
	"fmt"
	"sort"
)

type Movie struct {
	Name        string
	Genre       string
	IMDB        float64
	ReleaseYear int
}

type ByImdbAndReleaseYear []Movie

func (m ByImdbAndReleaseYear) Len() int { return len(m) }
func (m ByImdbAndReleaseYear) Less(i, j int) bool {
	if m[i].IMDB == m[j].IMDB {
		return m[i].ReleaseYear < m[j].ReleaseYear
	}

	return m[i].IMDB < m[j].IMDB
}
func (m ByImdbAndReleaseYear) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func main() {
	movies := []Movie{
		{"Inception", "Sci-Fi", 8.8, 2010},
		{"The Dark Knight", "Action", 8.8, 2008},
		{"Interstellar", "Sci-Fi", 8.6, 2014},
		{"The Godfather", "Crime", 9.2, 1972},
		{"Pulp Fiction", "Crime", 8.9, 1994},
	}

	sort.Sort(ByImdbAndReleaseYear(movies))

	fmt.Println(movies)
}
