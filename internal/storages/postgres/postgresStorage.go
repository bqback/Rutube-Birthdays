package postgresql

import (
	"birthdays/internal/config"

	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const step string = "storage"

func GetDBConnection(conf config.DatabaseConfig) (*sqlx.DB, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?application_name=%s&search_path=%s&connect_timeout=%d",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.AppName,
		conf.Schema,
		conf.ConnectionTimeout,
	)

	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
