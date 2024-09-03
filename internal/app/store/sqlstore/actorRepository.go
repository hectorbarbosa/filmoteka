package sqlstore

import (
	"database/sql"
	"filmoteka/internal/app/models"
)

type ActorRepository struct {
	store *Store
}

func (r *ActorRepository) Create(a models.Actor) (int, error) {
	if err := a.Validate(); err != nil {
		return 0, err
	}

	var id int
	if err := r.store.db.QueryRow(
		"INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id;",
		a.Name,
		a.Gender,
		a.BirthDate,
	).Scan(&id); err != nil {
		// sqlState := getSQLState(err)
		if err == sql.ErrNoRows {
			return 0, ErrResourceNotFound
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (r *ActorRepository) Find(id int) (models.Actor, error) {
	a := models.Actor{}
	if err := r.store.db.QueryRow(
		"SELECT id, name, gender, birth_date FROM actors WHERE id = $1;",
		id,
	).Scan(
		&a.Id,
		&a.Name,
		&a.Gender,
		&a.BirthDate,
	); err != nil {
		switch err {
		case sql.ErrNoRows:
			return models.Actor{}, ErrResourceNotFound
		default:
			return models.Actor{}, err
		}
	}

	return a, nil
}

func (r *ActorRepository) FindAll() ([]models.Actor, error) {
	a := &models.Actor{}
	actors := make([]models.Actor, 0)
	rows, err := r.store.db.Query(
		"SELECT id, name, gender, birth_date FROM actors;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&a.Id,
			&a.Name,
			&a.Gender,
			&a.BirthDate,
		)
		if err != nil {
			return nil, err
		}
		actors = append(actors, *a)
	}

	return actors, nil
}

func (r *ActorRepository) Delete(id int) error {
	result, err := r.store.db.Exec("DELETE FROM actors WHERE id=$1;", id)
	if err != nil {
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

func (r *ActorRepository) Update(a models.Actor) error {
	if err := a.Validate(); err != nil {
		return err
	}

	result, err := r.store.db.Exec(
		"UPDATE actors SET name=$1, gender=$2, birth_date=$3 WHERE id=$4;",
		a.Name,
		a.Gender,
		a.BirthDate,
		a.Id,
	)
	if err != nil {
		return err
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
