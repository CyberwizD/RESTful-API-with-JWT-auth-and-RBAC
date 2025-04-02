package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/handlers"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/utils"
	"github.com/golang-jwt/jwt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, store handlers.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(int64)

		_, err = store.GetUserById(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

func CreateJWT(secret []byte, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})

	tokenstring, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenstring, err
}

func CreateJWTAdmin(secret []byte, adminID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID":   strconv.Itoa(int(adminID)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})

	tokenstring, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenstring, err
}

func validateJWT(tokenstring string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteJSON(w, http.StatusUnauthorized, utils.ErrorResponse{
		Error: fmt.Errorf("permission Denied").Error(),
	})
}
