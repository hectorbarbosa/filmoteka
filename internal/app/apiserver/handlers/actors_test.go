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

func TestHandler_ActorCreate(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIActorRepository, a models.Actor)

	// testActor := models.Actor{
	// 	Id:        1,
	// 	Name:      "Name 1",
	// 	Gender:    "M",
	// 	BirthDate: "1995-01-12",
	// }

	tests := []struct {
		name                 string
		inputBody            string
		input                models.Actor
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name":"Name 1","gender":"M","birth_date":"1995-01-12"}`,
			input: models.Actor{
				Name:      "Name 1",
				Gender:    "M",
				BirthDate: "1995-01-12",
			},
			mockBehavior: func(r *mock_store.MockIActorRepository, a models.Actor) {
				r.EXPECT().Create(a).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            "",
			input:                models.Actor{},
			mockBehavior:         func(r *mock_store.MockIActorRepository, a models.Actor) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"EOF"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"name":"Name 1","gender":"M","birth_date":"1995-01-12"}`,
			input: models.Actor{
				Name:      "Name 1",
				Gender:    "M",
				BirthDate: "1995-01-12",
			},
			mockBehavior: func(r *mock_store.MockIActorRepository, a models.Actor) {
				r.EXPECT().Create(a).Return(0, errors.New(`something went wrong`))
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
			test.mockBehavior(actorRepo, test.input)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/actors", server.handleActorCreate()).Methods("POST")

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/actors",
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

func TestHandler_ActorFind(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIActorRepository, id int)

	testActor := models.Actor{
		Id:        1,
		Name:      "Name 1",
		Gender:    "M",
		BirthDate: "1995-01-12",
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
			mockBehavior: func(r *mock_store.MockIActorRepository, id int) {
				r.EXPECT().Find(id).Return(testActor, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"Name 1","gender":"M","birth_date":"1995-01-12"}`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			input:     1,
			mockBehavior: func(r *mock_store.MockIActorRepository, id int) {
				r.EXPECT().Find(id).Return(models.Actor{}, errors.New(`something went wrong`))
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
			test.mockBehavior(actorRepo, test.input)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/actors/{id}", server.handleActorFind()).Methods("GET")

			// Create Request
			w := httptest.NewRecorder()
			reqUrl := "/actors/" + strconv.Itoa(test.input)
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

func TestHandler_ActorFindAll(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIActorRepository)

	actors := []models.Actor{

		{
			Id:        1,
			Name:      "Name 1",
			Gender:    "M",
			BirthDate: "1995-01-12",
		},
		{
			Id:        2,
			Name:      "Name 2",
			Gender:    "F",
			BirthDate: "1995-02-12",
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
			mockBehavior: func(r *mock_store.MockIActorRepository) {
				r.EXPECT().FindAll().Return(actors, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":1,"name":"Name 1","gender":"M","birth_date":"1995-01-12"},{"id":2,"name":"Name 2","gender":"F","birth_date":"1995-02-12"}]`,
		},
		{
			name:      "Service Error",
			inputBody: ``,
			mockBehavior: func(r *mock_store.MockIActorRepository) {
				r.EXPECT().FindAll().Return(actors, errors.New(`something went wrong`))
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
			test.mockBehavior(actorRepo)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/actors", server.handleAllActors()).Methods("GET")

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/actors", bytes.NewBufferString(""))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}

func TestHandler_ActorDelete(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIActorRepository, id int)

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
			mockBehavior: func(r *mock_store.MockIActorRepository, id int) {
				r.EXPECT().Delete(id).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"result":true}`,
		},
		{
			name:      "Service Error",
			input:     1,
			inputBody: ``,
			mockBehavior: func(r *mock_store.MockIActorRepository, id int) {
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
			test.mockBehavior(actorRepo, test.input)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/actors/{id}", server.handleActorDelete()).Methods("DELETE")

			// Create Request
			w := httptest.NewRecorder()
			reqUrl := "/actors/" + strconv.Itoa(test.input)
			req := httptest.NewRequest("DELETE", reqUrl, bytes.NewBufferString(""))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}

func TestHandler_ActorUpdate(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_store.MockIActorRepository, a models.Actor)

	testActor := models.Actor{
		Id:        1,
		Name:      "Name 1",
		Gender:    "M",
		BirthDate: "1995-01-12",
	}

	tests := []struct {
		name                 string
		input                models.Actor
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			input:     testActor,
			inputBody: `{"name":"Name 1","gender":"M","birth_date":"1995-01-12"}`,
			mockBehavior: func(r *mock_store.MockIActorRepository, a models.Actor) {
				r.EXPECT().Update(a).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"Name 1","gender":"M","birth_date":"1995-01-12"}`,
		},
		{
			name:      "Service Error",
			input:     testActor,
			inputBody: `{"name":"Name 1","gender":"M","birth_date":"1995-01-12"}`,
			mockBehavior: func(r *mock_store.MockIActorRepository, a models.Actor) {
				r.EXPECT().Update(a).Return(errors.New(`something went wrong`))
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
			test.mockBehavior(actorRepo, test.input)
			store := mock_store.New(filmRepo, actorRepo)
			server := NewServer(store)

			// Init Endpoint
			router := mux.NewRouter()
			router.HandleFunc("/actors/{id}", server.handleActorUpdate()).Methods("PUT")

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/actors/1",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, strings.TrimRight(w.Body.String(), "\n"), test.expectedResponseBody)
		})
	}
}
