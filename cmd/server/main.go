package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	mux.HandleFunc("/api/v1/users", getUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id}", getUser).Methods("GET")
	mux.HandleFunc("/api/v1/users", createUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id}", updateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id}", deleteUser).Methods("DELETE")
	mux.HandleFunc("/api/v1/login", login).Methods("POST")
	mux.HandleFunc("/api/v1/logout", logout).Methods("POST")
	mux.HandleFunc("/api/v1/refresh", refresh).Methods("POST")
	mux.HandleFunc("/api/v1/roles", getRoles).Methods("GET")
	mux.HandleFunc("/api/v1/roles/{id}", getRole).Methods("GET")
	mux.HandleFunc("/api/v1/roles", createRole).Methods("POST")
	mux.HandleFunc("/api/v1/roles/{id}", updateRole).Methods("PUT")
	mux.HandleFunc("/api/v1/roles/{id}", deleteRole).Methods("DELETE")
	mux.HandleFunc("/api/v1/permissions", getPermissions).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{id}", getPermission).Methods("GET")
	mux.HandleFunc("/api/v1/permissions", createPermission).Methods("POST")
	mux.HandleFunc("/api/v1/permissions/{id}", updatePermission).Methods("PUT")
	mux.HandleFunc("/api/v1/permissions/{id}", deletePermission).Methods("DELETE")
	mux.HandleFunc("/api/v1/roles/{roleId}/permissions", addPermissionToRole).Methods("POST")
	mux.HandleFunc("/api/v1/roles/{roleId}/permissions/{permissionId}", removePermissionFromRole).Methods("DELETE")
	mux.HandleFunc("/api/v1/roles/{roleId}/users", addUserToRole).Methods("POST")
	mux.HandleFunc("/api/v1/roles/{roleId}/users/{userId}", removeUserFromRole).Methods("DELETE")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/users", addUserToPermission).Methods("POST")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/users/{userId}", removeUserFromPermission).Methods("DELETE")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles", addRoleToPermission).Methods("POST")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}", removeRoleFromPermission).Methods("DELETE")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/users", addUserToRolePermission).Methods("POST")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/users/{userId}", removeUserFromRolePermission).Methods("DELETE")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/users/{userId}/check", checkUserPermission).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/check", checkRolePermission).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/check", checkPermission).Methods("GET")
	mux.HandleFunc("/api/v1/roles/{roleId}/check", checkRole).Methods("GET")
	mux.HandleFunc("/api/v1/users/{userId}/check", checkUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/{userId}/roles", getUserRoles).Methods("GET")
	mux.HandleFunc("/api/v1/users/{userId}/permissions", getUserPermissions).Methods("GET")
	mux.HandleFunc("/api/v1/roles/{roleId}/users", getRoleUsers).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/users", getPermissionUsers).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles", getPermissionRoles).Methods("GET")
	mux.HandleFunc("/api/v1/roles/{roleId}/permissions", getRolePermissions).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/users", getRolePermissionUsers).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/check", checkRolePermission).Methods("GET")
	mux.HandleFunc("/api/v1/permissions/{permissionId}/roles/{roleId}/users/{userId}/check", checkUserRolePermission).Methods("GET")

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
