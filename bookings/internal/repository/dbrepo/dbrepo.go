package dbrepo

import (
	"database/sql"

	"github.com/MatthiasHeylenTM2021/golang-ppp/internal/config"
	"github.com/MatthiasHeylenTM2021/golang-ppp/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo implements repository.DatabaseRepo.
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// NewPostgresRepo implements repository.DatabaseRepo.for testing
func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
