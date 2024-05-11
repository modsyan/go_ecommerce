package entities

import "time"

type User struct {
	Id             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	HashedPassword string    `json:"hashed_password"`
	Role           UserRole  `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      string    `json:"updated_by"`
	DeletedAt      time.Time `json:"deleted_at"`
	DeletedBy      string    `json:"deleted_by"`
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)
