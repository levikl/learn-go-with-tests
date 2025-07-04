package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	const player string = "Pepper"

	t.Run(`GET "/players/%s" returns "3" after recording 3 wins via POST`, func(t *testing.T) {
		server := PlayerServer{NewInMemoryPlayerStore()}
		wins := 3

		for range wins {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		}

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		server := PlayerServer{NewInMemoryPlayerStore()}
		wins := 1000

		var wg sync.WaitGroup
		wg.Add(wins)

		for range wins {
			go func() {
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
				wg.Done()
			}()
		}
		wg.Wait()

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
	})
}
