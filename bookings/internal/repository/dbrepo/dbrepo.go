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

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

