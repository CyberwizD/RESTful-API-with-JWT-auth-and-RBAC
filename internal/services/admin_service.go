package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/config"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/auth"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/handlers"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/utils"
	"github.com/gorilla/mux"
)

var errAdminEmailRequired = errors.New("email is required")
var errAdminFirstNameRequired = errors.New("first name is required")
var errAdminLastNameRequired = errors.New("last name is required")
var errAdminPasswordRequired = errors.New("password is required")

type AdminService struct {
	admin_route handlers.User
}

func NewAdminService(a handlers.User) *AdminService {
	return &AdminService{
		admin_route: a,
	}
}

func (a *AdminService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/admin/register", a.handleAdminRegister).Methods("POST")
	r.HandleFunc("/admin/login", a.handleAdminLogin).Methods("POST")
}

func (a *AdminService) handleAdminRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *config.Admin

	err = json.Unmarshal(body, &payload)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateAdminPayload(payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error creating admin"})
		return
	}

	payload.Password = hashedPassword

	admin, err := a.admin_route.CreateAdmin(payload)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error creating admin"})
		return
	}

	token, err := createAndSetAuthCookieAdmin(admin.ID, w)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Erorr creating admin"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, token)
}

func validateAdminPayload(admin *config.Admin) error {
	if admin.Email == "" {
		return errAdminEmailRequired
	}

	if admin.FirstName == "" {
		return errAdminFirstNameRequired
	}

	if admin.LastName == "" {
		return errAdminLastNameRequired
	}

	if admin.Password == "" {
		return errAdminPasswordRequired
	}

	return nil
}

func createAndSetAuthCookieAdmin(adminID int64, w http.ResponseWriter) (string, error) {
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, adminID)

	if err != nil {
		return "Error creating auth cookie", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}

func (a *AdminService) handleAdminLogin(w http.ResponseWriter, r *http.Request) {

}
