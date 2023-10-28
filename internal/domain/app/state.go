package app

import (
	"context"
	"gojoke/internal/domain/joke"
)

type stateContextKey string

const stateCtxKey = stateContextKey("state")

type State struct {
	jokeService joke.Service
}

func NewState(jokeService joke.Service) State {
	return State{
		jokeService: jokeService,
	}
}

func (s *State) JokeService() *joke.Service {
	return &s.jokeService
}

func ContextWithState(state State, ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, stateCtxKey, state)
}

func StateFromContext(ctx context.Context) (*State, error) {
	state := ctx.Value(stateCtxKey).(State)

	return &state, nil
}
