package sqlstore

import (
	"database/sql"
	"filmoteka/internal/app/models"
	"strings"
)

type FilmRepository struct {
	store *Store
}

func (r *FilmRepository) Create(f models.Film) (int, error) {
	if err := f.Validate(); err != nil {
		// fmt.Println(ErrValidation.Error())
		return 0, ErrValidation
	}

	var id int
	if err := r.store.db.QueryRow(
		"INSERT INTO films (name, description, release_year, rating) VALUES ($1, $2, $3, $4) RETURNING id;",
		f.Name,
		f.Description,
		f.ReleaseYear,
		f.Rating,
	).Scan(&id); err != nil {
		// sqlState := getSQLState(err)
		if strings.Contains(err.Error(), "unique constraint") {
			return 0, ErrUniqueConstraints
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (r *FilmRepository) Find(id int) (models.Film, error) {
	f := models.Film{}
	if err := r.store.db.QueryRow(
		"SELECT id, name, description, release_year, rating FROM films WHERE id=$1",
		id,
	).Scan(
		&f.Id,
		&f.Name,
		&f.Description,
		&f.ReleaseYear,
		&f.Rating,
	); err != nil {
		switch err {
		case sql.ErrNoRows:
			return models.Film{}, ErrResourceNotFound
		default:
			return models.Film{}, err
		}
	}

	return f, nil
}

func (r *FilmRepository) FindAll() ([]models.Film, error) {
	f := &models.Film{}
	films := make([]models.Film, 0)
	rows, err := r.store.db.Query(
		"SELECT id, name, description, release_year, rating FROM films;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&f.Id,
			&f.Name,
			&f.Description,
			&f.ReleaseYear,
			&f.Rating,
		)
		if err != nil {
			return nil, err
		}
		films = append(films, *f)
	}

	return films, nil
}

func (r *FilmRepository) Delete(id int) error {
	result, err := r.store.db.Exec("DELETE FROM films WHERE id=$1;", id)
	if err != nil {
		// fmt.Println(err.Error())
		return err
	}

	deletedRows, err := result.RowsAffected()
	// fmt.Println("deleted: ", deletedRows)
	if err != nil {
		return err
	}
	if deletedRows == 0 {
		return ErrResourceNotFound
	}
	return nil
}

func (r *FilmRepository) Update(f models.Film) error {
	if err := f.Validate(); err != nil {
		return err
	}

	result, err := r.store.db.Exec(
		"UPDATE films SET name=$1, description=$2, release_year=$3, rating=$4 WHERE id=$5;",
		f.Name,
		f.Description,
		f.ReleaseYear,
		f.Rating,
		f.Id,
	)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return ErrUniqueConstraints
		} else {
			return err
		}
	}

	updatedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if updatedRows == 0 {
		return ErrResourceNotFound
	}

	return nil
}
