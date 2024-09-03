package store

import (
	"filmoteka/internal/app/models"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type IFilmRepository interface {
	Create(models.Film) (int, error)
	Find(int) (models.Film, error)
	FindAll() ([]models.Film, error)
	Delete(id int) error
	Update(models.Film) error
}

type IActorRepository interface {
	Create(models.Actor) (int, error)
	Find(int) (models.Actor, error)
	FindAll() ([]models.Actor, error)
	Delete(id int) error
	Update(models.Actor) error
}
