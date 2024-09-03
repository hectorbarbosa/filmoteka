package sqlstore

import (
	"database/sql"
	"filmoteka/internal/app/models"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestActor_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	testActor := models.Actor{
		Name:      "Name 1",
		Gender:    "M",
		BirthDate: "1995-01-12",
	}

	type args struct {
		actor models.Actor
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
				actor: testActor,
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery(
					"INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id;",
				).WithArgs(
					args.actor.Name,
					args.actor.Gender,
					args.actor.BirthDate,
				).WillReturnRows(rows)
			},
		},
		{
			name: "Failed empty field",
			input: args{
				actor: models.Actor{
					Gender:    "M",
					BirthDate: "1995-01-12",
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

			got, err := r.ActorRepo().Create(tt.input.actor)
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

func TestActor_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	tests := []struct {
		name    string
		mock    func()
		want    []models.Actor
		wantErr bool
	}{
		{
			name: "Regular Select",
			mock: func() {
				rows := sqlmock.NewRows([]string{
					"id", "name", "gender", "birth_date",
				}).
					AddRow(1, "Actor1", "M", "1980-01-12").
					AddRow(2, "Actor2", "M", "1990-02-20").
					AddRow(3, "Actor3", "F", "1990-02-20")

				mock.ExpectQuery(
					"SELECT id, name, gender, birth_date FROM actors;",
				).WithArgs().WillReturnRows(rows)
			},
			want: []models.Actor{
				{Id: 1, Name: "Actor1", Gender: "M", BirthDate: "1980-01-12"},
				{Id: 2, Name: "Actor2", Gender: "M", BirthDate: "1990-02-20"},
				{Id: 3, Name: "Actor3", Gender: "F", BirthDate: "1990-02-20"},
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows(
					[]string{"id", "name", "gender", "birth_date"})

				mock.ExpectQuery(
					"SELECT id, name, gender, birth_date FROM actors;",
				).WithArgs().WillReturnRows(rows)
			},
			want: []models.Actor{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.ActorRepo().FindAll()
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

func TestActorFind(t *testing.T) {
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
		want    models.Actor
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "gender", "birth_date",
				}).AddRow(1, "Name 1", "M", "1980-01-01")
				mock.ExpectQuery( // regexp.QuoteMeta( -- also works
					"SELECT id, name, gender, birth_date FROM actors WHERE id = $1;",
				).WithArgs(args.id).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want: models.Actor{
				Id: 1, Name: "Name 1", Gender: "M", BirthDate: "1980-01-01",
			},
		},
		{
			name:  "NotFound",
			input: args{id: 700},
			mock: func(args args) {
				// regexp.QuoteMeta -- also works
				mock.ExpectQuery( // regexp.QuoteMeta(
					"SELECT id, name, gender, birth_date FROM actors WHERE id = $1;",
				).WithArgs(args.id).WillReturnError(ErrResourceNotFound)
			},
			// want:    &models.Film{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.ActorRepo().Find(tt.input.id)
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

func TestActorDelete(t *testing.T) {
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
				mock.ExpectExec("DELETE FROM actors WHERE id=$1;").
					WithArgs(args.id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "Not Found",
			input: args{
				id: 404,
			},
			mock: func(args args) {
				mock.ExpectExec("DELETE FROM actors WHERE id=$1;").
					WithArgs(args.id).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			err := r.ActorRepo().Delete(tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestActorUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	updatedActor := models.Actor{
		Id:        1,
		Name:      "Name 1",
		Gender:    "M",
		BirthDate: "1995-01-12",
	}

	type args struct {
		id    int
		actor models.Actor
	}
	tests := []struct {
		name    string
		mock    func(args args, a *models.Actor)
		input   args
		want    *models.Actor
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id:    1,
				actor: updatedActor,
			},
			mock: func(args args, a *models.Actor) {
				mock.ExpectExec(
					"UPDATE actors SET name=$1, gender=$2, birth_date=$3 WHERE id=$4;",
				).WithArgs(
					args.actor.Name,
					args.actor.Gender,
					args.actor.BirthDate,
					args.id,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			want: &updatedActor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			err := r.ActorRepo().Update(tt.input.actor)
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
