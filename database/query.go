package database

// Predicate is a string that acts as a condition in the where clause
type Predicate string

type query struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func NewQueryBuilder() *query {
	return new(query)
}

func First() (*Entry, error) {
	ret := db.First(Entry)

	return ret, res.Error
}
