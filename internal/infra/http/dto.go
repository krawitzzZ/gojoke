package http

import (
	"gojoke/internal/domain/joke"
	"time"

	"github.com/google/uuid"
)

type JokeDto struct {
	ID            uuid.UUID `json:"id"`
	Body          string    `json:"body"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

func (j JokeDto) ToModel() joke.Model {
	return joke.Model{
		ID:            j.ID,
		Body:          j.Body,
		CreatedAt:     j.CreatedAt,
		LastUpdatedAt: j.LastUpdatedAt,
	}
}

type JokeQueryDto struct {
	Items       []JokeDto `json:"items"`
	CurrentPage int       `json:"currentPage"`
	TotalPages  int       `json:"totalPages"`
	TotalCount  int       `json:"totalCount"`
}

func (jqd JokeQueryDto) ToModel(query string) joke.ModelQuery {
	jokes := make([]joke.Model, len(jqd.Items))
	for i, joke := range jqd.Items {
		jokes[i] = joke.ToModel()
	}

	return joke.NewModelQuery(jokes, query, jqd.CurrentPage, jqd.TotalPages)
}

type NewJokeDto struct {
	Body string `json:"body"`
}
