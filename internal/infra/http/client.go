package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gojoke/internal/infra"
	"net/http"
	"net/url"
)

const baseURL = "http://localhost:4343"

func GetStatus() bool {
	resp, err := http.Get(baseURL + "/api/status")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func getRandomJoke() (JokeDto, error) {
	var result JokeDto

	resp, err := http.Get(baseURL + "/api/jokes/random")
	if err != nil {
		inner := fmt.Errorf("get random joke request failed: %w", err)
		return result, infra.NewHttpError(inner)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		inner := fmt.Errorf("failed to decode get random joke request: %w", err)
		return result, infra.NewDecodingError(inner)
	}

	return result, nil
}

func queryJokes(query string) (JokeQueryDto, error) {
	var result JokeQueryDto

	resp, err := http.Get(baseURL + "/api/jokes?b=" + url.QueryEscape(query))
	if err != nil {
		inner := fmt.Errorf("query jokes request failed: %w", err)
		return result, infra.NewHttpError(inner)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		inner := fmt.Errorf("failed to decode query jokes request: %w", err)
		return result, infra.NewDecodingError(inner)
	}

	return result, nil
}

func createNewJoke(body string) (JokeDto, error) {
	var result JokeDto
	jokeDto := JokeDto{Body: body}

	jsonBody, err := json.Marshal(jokeDto)
	if err != nil {
		inner := fmt.Errorf("failed to encode joke dto: %w", err)
		return result, infra.NewEncodingError(inner)
	}

	resp, err := http.Post(baseURL+"/api/jokes", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		inner := fmt.Errorf("create new joke request failed: %w", err)
		return result, infra.NewHttpError(inner)
	}
	defer resp.Body.Close()

	return result, nil
}
