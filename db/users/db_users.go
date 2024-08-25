package db_users

import (
	"context"

	"scketch_finance_go/structs"
	"scketch_finance_go/db"

	"github.com/google/uuid"
)

func GetUser(db *db.Database, id uuid.UUID) (structs.User, error) {
	var user structs.User
	query := "SELECT id, name FROM users WHERE id = $1"
	err := db.Pool.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name)
	return user, err
}