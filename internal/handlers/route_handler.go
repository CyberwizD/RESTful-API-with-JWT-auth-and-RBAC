package handlers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/config"
)

type Storage struct {
	db *sql.DB
}

type User interface {
	// User
	CreateUser(u *config.User) (*config.User, error)
	GetUserById(id int64) (*config.User, error)
	DeleteUser(del *config.User) (*config.User, error)

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
	var exists bool

	// Check if the email already exists in the database
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", u.Email).Scan(&exists)

	if err != nil {
		log.Printf("Error checking user existence: %v", err)
		return nil, err
	}

	if exists {
		return nil, errors.New("user with this email already exists")
	}

	// Insert the new user if the email does not exist
	rows, err := s.db.Exec(`
		INSERT INTO users (email, firstName, lastName, password, role)
		VALUES (?, ?, ?, ?, 'user')
	`, u.Email, u.FirstName, u.LastName, u.Password)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	u.ID = id
	u.Role = "user"

	return u, nil
}

func (s *Storage) DeleteUser(del *config.User) (*config.User, error) {
	// Check if the user exists before attempting to delete
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ? AND role = 'user')", del.Email).Scan(&exists)

	if err != nil {
		log.Printf("Error checking user existence: %v", err)
		return nil, err
	}

	if !exists {
		return nil, errors.New("user does not exist")
	}

	// Delete the user from the database
	_, err = s.db.Exec("DELETE FROM users WHERE email = ? AND role = 'user'", del.Email)

	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, err
	}

	return del, nil
}

// GetUserByEmail retrieves an admin user by their email.
func (s *Storage) GetUserById(id int64) (*config.User, error) {
	var user config.User
	query := "SELECT id, email, firstName, lastName, role FROM users WHERE id = ? AND role = 'user'"
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Storage) CreateAdmin(a *config.Admin) (*config.Admin, error) {
	rows, err := s.db.Exec(`
		INSERT INTO admins (email, firstName, lastName, password, role)
		VALUES (?, ?, ?, ?, 'admin')
	`, a.Email, a.FirstName, a.LastName, a.Password)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	a.ID = id
	a.Role = "admin"

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
