package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"
	wins := 3

	for range wins {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", wins},
		}
		assertLeague(t, got, want)
	})
}

// todo: this is broken now
// func TestConcurrency(t *testing.T) {
// 	database, cleanDatabase := createTempFile(t, `[]`)
// 	defer cleanDatabase()
//
// 	store, err := NewFileSystemPlayerStore(database)
// 	assertNoError(t, err)
//
// 	server := NewPlayerServer(store)
// 	player := "Vegeta"
// 	wins := 9001
//
// 	var wg sync.WaitGroup
// 	wg.Add(wins)
//
// 	for range wins {
// 		go func() {
// 			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
//
// 	t.Run("get score", func(t *testing.T) {
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, newGetScoreRequest(player))
// 		assertStatus(t, response.Code, http.StatusOK)
//
// 		assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
// 	})
//
// 	t.Run("get league", func(t *testing.T) {
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, newLeagueRequest())
// 		assertStatus(t, response.Code, http.StatusOK)
//
// 		got := getLeagueFromResponse(t, response.Body)
// 		want := []Player{
// 			{"Vegeta", wins},
// 		}
// 		assertLeague(t, got, want)
// 	})
// }
