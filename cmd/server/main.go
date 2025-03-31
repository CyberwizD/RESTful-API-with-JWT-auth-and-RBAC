package main

import (
	"log"

	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/config"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/api"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/handlers"
	"github.com/CyberwizD/RESTful-API-with-JWT-auth-and-RBAC/internal/storage"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := storage.NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()

	if err != nil {
		log.Fatal(err)
	}

	route := handlers.NewStorage(db)

	apiServer := api.NewAPIServer(":8080", route)
	apiServer.Serve()
}
