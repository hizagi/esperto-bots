package entities

import "time"

type User struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Password  string     `json:"password"`
	Document  string     `json:"document"`
	BirthDate string     `json:"birth_date"`
	Roles     []*Role    `json:"roles"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
