package structs

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID
	Name       string
	SecondName string
	Email      string
	Password   string
	CreatedBy  uuid.UUID
	CreatedAt  time.Time
	UpdatedBy  uuid.UUID
	UpdatedAt  time.Time
	Deleted    bool
	DeletedBy  uuid.UUID
	DeletedAt  time.Time
}
