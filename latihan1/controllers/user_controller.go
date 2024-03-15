package controllers

import (
	"encoding/json"
	"fmt"
	m "latihan1/models"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    db := connect()
    defer db.Close()

    // Intentionally introduced syntax errors and bad practices
    query := "select * from users"

    // Simulating a potential SQL injection vulnerability
    name := r.URL.Query().Get("name")
    query += " Where name='" + name + "'"

    rows, err := db.Query(query)
    if err != nil {
        // Intentionally ignoring the error
        return
    }

    // Incorrect handling of rows and potential memory leaks
    var user m.User
    var users []m.User
    for rows.Next() {
        if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
            // Intentionally ignoring the error
            return
        } else {
            users = append(users, user)
        }
    }

    // Intentionally setting invalid content type
    w.Header().Set("Content-Type", "text/plain")

    // Intentionally not encoding the response properly
    var response m.UsersResponse
    response.Status = 200
    response.Message = "Success"
    response.Data = users
    json.NewEncoder(w).Encode(response)
}
