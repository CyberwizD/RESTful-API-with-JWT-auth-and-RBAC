package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/config"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/auth"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/handlers"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/utils"
	"github.com/gorilla/mux"
)

var errEmailRequired = errors.New("email is required")
var errFirstNameRequired = errors.New("first name is required")
var errLastNameRequired = errors.New("last name is required")
var errPasswordRequired = errors.New("password is required")

type UserService struct {
	user_route handlers.User
}

func NewUserService(s handlers.User) *UserService {
	return &UserService{
		user_route: s,
	}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", s.handleUserLogin).Methods("POST")
	r.HandleFunc("/users/{id}", s.handlerUserId).Methods("GET")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Use RegisterRequest instead of User directly
	var registerRequest *config.RegisterRequest
	err = json.Unmarshal(body, &registerRequest)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	// Convert RegisterRequest to User
	payload := &config.User{
		Email:     registerRequest.Email,
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Password:  registerRequest.Password,
	}

	// Validate user input
	if err := validateUserPayload(payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error creating user"})
		return
	}

	payload.Password = hashedPassword

	// Create user in storage
	user, err := s.user_route.CreateUser(payload)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error creating user"})
		return
	}

	// Create JWT token
	token, err := createAndSetAuthCookie(user.ID, w)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error creating user"})
		return
	}

	// Response: Success!
	utils.WriteJSON(w, http.StatusCreated, utils.SuccessResponse{
		Message: "User registered successfully",
		Token:   token,
	})
}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {

}

func validateUserPayload(user *config.User) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func (s *UserService) handlerUserId(w http.ResponseWriter, r *http.Request) {
	userVar := mux.Vars(r)
	IdStr := userVar["id"]

	userId, err := strconv.ParseInt(IdStr, 10, 64)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "User not found"})
		return
	}

	user, err := s.user_route.GetUserById(userId)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "User not found"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func createAndSetAuthCookie(userID int64, w http.ResponseWriter) (string, error) {
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, userID)

	if err != nil {
		return "Error creating auth cookie", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
