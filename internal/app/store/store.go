package store

type IStore interface {
	FilmRepo() IFilmRepository
	ActorRepo() IActorRepository
}
