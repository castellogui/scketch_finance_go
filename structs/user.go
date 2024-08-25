package structs

import (
    "database/sql"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID
	Name       string
	SecondName string
	Email      string
	Password   string
	CreatedBy  uuid.UUID
	CreatedAt  sql.NullTime
	UpdatedBy  uuid.UUID
	UpdatedAt  sql.NullTime
	Deleted    bool
	DeletedBy  uuid.UUID
	DeletedAt  sql.NullTime
}
