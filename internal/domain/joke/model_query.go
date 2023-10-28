package joke

import "errors"

type ModelQuery struct {
	Jokes       []Model
	query       string
	currentPage int
	totalPages  int
}

func NewModelQuery(jokes []Model, query string, currentPage, totalPages int) ModelQuery {
	return ModelQuery{
		Jokes:       jokes,
		query:       query,
		currentPage: currentPage,
		totalPages:  totalPages,
	}
}

func (jq *ModelQuery) Query() string {
	return jq.query
}

func (jq *ModelQuery) HasPrevPage() bool {
	return jq.totalPages > 1 && jq.currentPage > 1
}

func (jq *ModelQuery) HasNextPage() bool {
	return jq.totalPages > 1 && jq.currentPage < jq.totalPages
}

func (jq *ModelQuery) GetPrevPage() (int, error) {
	prevPage := jq.currentPage - 1
	if prevPage <= 0 {
		return jq.currentPage, errors.New("no previous page")
	}

	return prevPage, nil
}

func (jq *ModelQuery) GetNextPage() (int, error) {
	nextPage := jq.currentPage + 1
	if nextPage > jq.totalPages {
		return jq.currentPage, errors.New("no next page")
	}

	return nextPage, nil
}
