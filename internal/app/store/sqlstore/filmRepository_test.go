package sqlstore

import (
	"database/sql"
	"filmoteka/internal/app/models"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFilm_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	testFilm := models.Film{
		Name:        "Test Film 1",
		Description: "Descr 1",
		ReleaseYear: 2015,
		Rating:      6.7,
	}
	type args struct {
		film models.Film
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "RegularInsert",
			input: args{
				film: testFilm,
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery(
					"INSERT INTO films (name, description, release_year, rating) VALUES ($1, $2, $3, $4) RETURNING id;",
				).WithArgs(
					args.film.Name,
					args.film.Description,
					args.film.ReleaseYear,
					args.film.Rating,
				).WillReturnRows(rows)
			},
		},
		{
			name: "Failed empty field",
			input: args{
				film: models.Film{
					Description: "Descr 1",
					ReleaseYear: 2015,
					Rating:      8,
				},
			},
			mock: func(args args, id int) {
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.FilmRepo().Create(tt.input.film)
			// fmt.Println("got: ", got)
			// fmt.Println("err: ", err)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				// fmt.Println("want: ", tt.want)
				// fmt.Println("got: ", got)

				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestFilm_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	tests := []struct {
		name    string
		mock    func()
		want    []models.Film
		wantErr bool
	}{
		{
			name: "Regular Select",
			mock: func() {
				rows := sqlmock.NewRows([]string{
					"id", "name", "description", "release_year", "rating",
				}).
					AddRow(1, "film1", "description1", 2000, 10).
					AddRow(2, "film2", "description2", 2001, 4).
					AddRow(3, "film3", "description3", 2002, 5)

				mock.ExpectQuery("SELECT id, name, description, release_year, rating FROM films;").
					WithArgs().WillReturnRows(rows)
			},
			want: []models.Film{
				{Id: 1, Name: "film1", Description: "description1", ReleaseYear: 2000, Rating: 10},
				{Id: 2, Name: "film2", Description: "description2", ReleaseYear: 2001, Rating: 4},
				{Id: 3, Name: "film3", Description: "description3", ReleaseYear: 2002, Rating: 5},
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "name", "description", "release_year, rating"})

				mock.ExpectQuery("SELECT id, name, description, release_year, rating FROM films;").
					WithArgs().WillReturnRows(rows)
			},
			want: []models.Film{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.FilmRepo().FindAll()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				// fmt.Println("want: ", tt.want)
				// fmt.Println("got: ", got)
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestFind(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		mock    func(args args)
		input   args
		want    models.Film
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "description", "release_year", "rating",
				}).AddRow(1, "film1", "description1", 2000, 10)
				mock.ExpectQuery( // regexp.QuoteMeta( -- also works
					"SELECT id, name, description, release_year, rating FROM films WHERE id=$1",
				).WithArgs(args.id).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want: models.Film{
				Id: 1, Name: "film1", Description: "description1", ReleaseYear: 2000, Rating: 10,
			},
		},
		{
			name:  "NotFound",
			input: args{id: 700},
			mock: func(args args) {
				// regexp.QuoteMeta -- also works
				mock.ExpectQuery( // regexp.QuoteMeta(
					"SELECT id, name, description, release_year, rating FROM films WHERE id=$1",
				).WithArgs(args.id).WillReturnError(ErrResourceNotFound)
			},
			// want:    &models.Film{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.FilmRepo().Find(tt.input.id)
			// fmt.Println("got: ", got)
			// fmt.Println("err: ", err)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				// fmt.Println("want: ", tt.want)
				// fmt.Println("got: ", got)
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		mock    func(args args)
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				mock.ExpectExec("DELETE FROM films WHERE id=$1;").
					WithArgs(args.id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "Not Found",
			input: args{
				id: 404,
			},
			mock: func(args args) {
				mock.ExpectExec("DELETE FROM films WHERE id=$1;").
					WithArgs(args.id).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			err := r.FilmRepo().Delete(tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	updatedFilm := models.Film{
		Id:          1,
		Name:        "Updated Film",
		Description: "Updated Description",
		ReleaseYear: 2015,
		Rating:      6.7,
	}

	type args struct {
		id   int
		film models.Film
	}
	tests := []struct {
		name    string
		mock    func(args args, film *models.Film)
		input   args
		want    *models.Film
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id:   1,
				film: updatedFilm,
			},
			mock: func(args args, film *models.Film) {
				mock.ExpectExec("UPDATE films SET name=$1, description=$2, release_year=$3, rating=$4 WHERE id=$5;").
					WithArgs(
						args.film.Name,
						args.film.Description,
						args.film.ReleaseYear,
						args.film.Rating,
						args.id,
					).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			want: &updatedFilm,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			err := r.FilmRepo().Update(tt.input.film)
			// fmt.Println("got: ", got)
			// fmt.Println("err: ", err)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
