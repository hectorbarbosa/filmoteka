package sqlstore

import (
	"database/sql"

	"filmoteka/internal/app/store"
)

type Store struct {
	db              *sql.DB
	filmRepository  *FilmRepository
	actorRepository *ActorRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) FilmRepo() store.IFilmRepository {
	if s.filmRepository != nil {
		return s.filmRepository
	}

	s.filmRepository = &FilmRepository{
		store: s,
	}

	return s.filmRepository
}

func (s *Store) ActorRepo() store.IActorRepository {
	if s.actorRepository != nil {
		return s.actorRepository
	}

	s.actorRepository = &ActorRepository{
		store: s,
	}

	return s.actorRepository
}
