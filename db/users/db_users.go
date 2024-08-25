package db_users

import (
	"context"

	"scketch_finance_go/db"
	"scketch_finance_go/structs"

	"github.com/google/uuid"
)

func GetUser(db *db.Database, id uuid.UUID) (structs.User, error) {
	var user structs.User
	query := "SELECT * FROM users WHERE id = $1"
	err := db.Pool.QueryRow(context.Background(), query, id).Scan(
		&user.Id,
		&user.Name,
		&user.SecondName,
		&user.Email,
		&user.Password,
		&user.CreatedBy,
		&user.CreatedAt,
		&user.UpdatedBy,
		&user.UpdatedAt,
		&user.Deleted,
		&user.DeletedBy,
		&user.DeletedAt,
	)
	return user, err
}
