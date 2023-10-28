package main

import (
	"gojoke/cmd/cli"
	"gojoke/internal/domain/app"
	"gojoke/internal/domain/joke"
	"gojoke/internal/infra/http"
)

func main() {
	jokeRepo := http.JokeRepository{}
	jokeService := joke.NewService(&jokeRepo)
	appState := app.NewState(jokeService)

	cli.Execute(appState)
}
