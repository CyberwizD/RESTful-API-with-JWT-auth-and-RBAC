package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL!")

	return &MySQLStorage{
		db: db,
	}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	// initialize the tables
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createAdminsTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQLStorage) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *MySQLStorage) createAdminsTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS admins (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *MySQLStorage) Close() error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	return nil
}
