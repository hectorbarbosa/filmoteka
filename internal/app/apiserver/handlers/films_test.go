package handlers

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"filmoteka/internal/app/models"
	"filmoteka/internal/app/store/mock_store"
)

func TestHandler_FilmCreate(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIFilmRepository, film models.Film)

	// testFilm := models.Film{
	// 	Name:        "Test Name",
	// 	Description: "Desc1",
	// 	ReleaseYear: 2002,
	// 	Rating:      7.5,
	// }

	tests := []struct {
		name                 string
		inputBody            string
		inputFilm            models.Film
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name":"Test Name","description":"Desc1","release_year":2002,"rating":7.5}`,
			inputFilm: models.Film{
				Name:        "Test Name",
				Description: "Desc1",
				ReleaseYear: 2002,
				Rating:      7.5,
			},
			mockBehavior: func(r *mock_store.MockIFilmRepository, film models.Film) {
				r.EXPECT().Create(film).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            "",
			inputFilm:            models.Film{},
			mockBehavior:         func(r *mock_store.MockIFilmRepository, film models.Film) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"EOF"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"name": "Test Name", "description": "Desc1", "release_year": 2002, "rating": 7.5}`,
			inputFilm: models.Film{
				Name:        "Test Name",
				Description: "Desc1",
				ReleaseYear: 2002,
				Rating:      7.5,
			},
			mockBehavior: func(r *mock_store.MockIFilmRepository, film models.Film) {
				r.EXPECT().Create(film).Return(0, errors.New(`something went wrong`))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			filmRepo := mock_store.NewMockIFilmRepository(c)
			actorRepo := mock_store.NewMockIActorRepository(c)
			test.mockBehavior(filmRepo, test.inputFilm)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/films", server.handleFilmCreate()).Methods("POST")

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/films",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			// fmt.Println("Body :", w.Body.String())
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.Trim(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}

func TestHandler_FilmFind(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIFilmRepository, id int)

	testFilm := models.Film{
		Id:          1,
		Name:        "Test Name",
		Description: "Desc1",
		ReleaseYear: 2002,
		Rating:      7.5,
	}

	tests := []struct {
		name                 string
		inputBody            string
		input                int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: ``,
			input:     1,
			mockBehavior: func(r *mock_store.MockIFilmRepository, id int) {
				r.EXPECT().Find(id).Return(testFilm, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"Test Name","description":"Desc1","release_year":2002,"rating":7.5}`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			input:     1,
			mockBehavior: func(r *mock_store.MockIFilmRepository, id int) {
				r.EXPECT().Find(id).Return(models.Film{}, errors.New(`something went wrong`))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			filmRepo := mock_store.NewMockIFilmRepository(c)
			actorRepo := mock_store.NewMockIActorRepository(c)
			test.mockBehavior(filmRepo, test.input)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/films/{id}", server.handleFilmFind()).Methods("GET")

			// Create Request
			w := httptest.NewRecorder()
			reqUrl := "/films/" + strconv.Itoa(test.input)
			req := httptest.NewRequest("GET", reqUrl, bytes.NewBufferString(""))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			// fmt.Println("Body :", w.Body.String())
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}

func TestHandler_FilmFindAll(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIFilmRepository)

	films := []models.Film{

		{
			Id:          1,
			Name:        "Test Name",
			Description: "Desc1",
			ReleaseYear: 2002,
			Rating:      7.5,
		},
		{
			Id:          2,
			Name:        "Test Name 2",
			Description: "Desc2",
			ReleaseYear: 2004,
			Rating:      8.5,
		},
	}

	tests := []struct {
		name                 string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: ``,
			mockBehavior: func(r *mock_store.MockIFilmRepository) {
				r.EXPECT().FindAll().Return(films, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":1,"name":"Test Name","description":"Desc1","release_year":2002,"rating":7.5},{"id":2,"name":"Test Name 2","description":"Desc2","release_year":2004,"rating":8.5}]`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			mockBehavior: func(r *mock_store.MockIFilmRepository) {
				r.EXPECT().FindAll().Return(films, errors.New(`something went wrong`))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			filmRepo := mock_store.NewMockIFilmRepository(c)
			actorRepo := mock_store.NewMockIActorRepository(c)
			test.mockBehavior(filmRepo)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/films", server.handleAllFilms()).Methods("GET")

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/films", bytes.NewBufferString(""))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}

func TestHandler_FilmDelete(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIFilmRepository, id int)

	tests := []struct {
		name                 string
		input                int
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			input:     1,
			inputBody: ``,
			mockBehavior: func(r *mock_store.MockIFilmRepository, id int) {
				r.EXPECT().Delete(id).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"result":true}`,
		},
		{
			name:      "Service Error",
			input:     1,
			inputBody: ``,
			mockBehavior: func(r *mock_store.MockIFilmRepository, id int) {
				r.EXPECT().Delete(id).Return(errors.New(`something went wrong`))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			filmRepo := mock_store.NewMockIFilmRepository(c)
			actorRepo := mock_store.NewMockIActorRepository(c)
			test.mockBehavior(filmRepo, test.input)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/films/{id}", server.handleFilmDelete()).Methods("DELETE")

			// Create Request
			w := httptest.NewRecorder()
			reqUrl := "/films/" + strconv.Itoa(test.input)
			req := httptest.NewRequest("DELETE", reqUrl, bytes.NewBufferString(""))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}

func TestHandler_FilmUpdate(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIFilmRepository, film models.Film)

	testFilm := models.Film{
		Id:          1,
		Name:        "Test Name",
		Description: "Desc1",
		ReleaseYear: 2002,
		Rating:      7.5,
	}

	tests := []struct {
		name                 string
		inputFilm            models.Film
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputFilm: testFilm,
			inputBody: `{"name":"Test Name","description":"Desc1","release_year":2002,"rating":7.5}`,
			mockBehavior: func(r *mock_store.MockIFilmRepository, film models.Film) {
				r.EXPECT().Update(film).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"Test Name","description":"Desc1","release_year":2002,"rating":7.5}`,
		},
		{
			name:      "Service Error",
			inputFilm: testFilm,
			inputBody: `{"name":"Test Name","description":"Desc1","release_year":2002,"rating":7.5}`,
			mockBehavior: func(r *mock_store.MockIFilmRepository, film models.Film) {
				r.EXPECT().Update(film).Return(errors.New(`something went wrong`))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			filmRepo := mock_store.NewMockIFilmRepository(c)
			actorRepo := mock_store.NewMockIActorRepository(c)
			test.mockBehavior(filmRepo, test.inputFilm)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/films/{id}", server.handleFilmUpdate()).Methods("PUT")

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/films/1",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}
