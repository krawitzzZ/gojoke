package http

import (
	"gojoke/internal/domain/app"
	"gojoke/internal/domain/joke"
)

type JokeRepository struct{}

func (r *JokeRepository) Random() (joke.Model, error) {
	jokeDto, err := getRandomJoke()
	if err != nil {
		return joke.Model{}, app.NewInternalError(err)
	}

	return jokeDto.ToModel(), err
}

func (r *JokeRepository) Query(query string) (joke.ModelQuery, error) {
	jokesQueryDto, err := queryJokes(query)
	if err != nil {
		return joke.ModelQuery{}, app.NewInternalError(err)
	}

	return jokesQueryDto.ToModel(query), err
}
