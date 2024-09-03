package apiserver

import (
	"database/sql"
	"net/http"

	"filmoteka/internal/app/apiserver/handlers"
	"filmoteka/internal/app/store/sqlstore"

	_ "github.com/lib/pq"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := handlers.NewServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
