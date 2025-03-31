package handlers

import (
	"database/sql"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/config"
)

type Storage struct {
	db *sql.DB
}

type User interface {
	// User
	CreateUser(u *config.User) (*config.User, error)
	GetUserByEmail(email string) (*config.User, error)

	// Admin
	CreateAdmin(a *config.Admin) (*config.Admin, error)
	GetAdminByEmail(email string) (*config.Admin, error)
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser(u *config.User) (*config.User, error) {
	rows, err := s.db.Exec(`
		INSERT INTO users (email, firstName, lastName, password)
		VALUES (?, ?, ?, ?)
	`, u.Email, u.FirstName, u.LastName, u.Password)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	u.ID = id

	return u, nil
}

// GetUserByEmail retrieves an admin user by their email.
func (s *Storage) GetUserByEmail(email string) (*config.User, error) {
	var user config.User
	query := "SELECT id, email, role FROM users WHERE email = ? AND role = 'user'"
	err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) CreateAdmin(a *config.Admin) (*config.Admin, error) {
	rows, err := s.db.Exec(`
		INSERT INTO admins (email, firstName, lastName, password)
		VALUES (?, ?, ?, ?)
	`, a.Email, a.FirstName, a.LastName, a.Password)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	a.ID = id

	return a, nil
}

// GetAdminByEmail retrieves an admin user by their email.
func (s *Storage) GetAdminByEmail(email string) (*config.Admin, error) {
	var admin config.Admin
	query := "SELECT id, email, role FROM users WHERE email = ? AND role = 'admin'"
	err := s.db.QueryRow(query, email).Scan(&admin.ID, &admin.Email, &admin.Role)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
