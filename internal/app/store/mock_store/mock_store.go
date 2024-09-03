package mock_store

import (
	"filmoteka/internal/app/store"
)

type MockStore struct {
	filmRepository  *MockIFilmRepository
	actorRepository *MockIActorRepository
}

func New(
	filmRepo *MockIFilmRepository,
	actorRepo *MockIActorRepository,
) *MockStore {
	return &MockStore{
		filmRepository:  filmRepo,
		actorRepository: actorRepo,
	}
}

func (s *MockStore) FilmRepo() store.IFilmRepository {
	return s.filmRepository
}

func (s *MockStore) ActorRepo() store.IActorRepository {
	return s.actorRepository
}
