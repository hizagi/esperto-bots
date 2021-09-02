package entities

import "time"

type Permission struct {
	ID           string     `json:"id"`
	ResourceName string     `json:"resource_name"`
	CanCreate    bool       `json:"can_create"`
	CanRead      bool       `json:"can_read"`
	CanUpdate    bool       `json:"can_update"`
	CanDelete    bool       `json:"can_delete"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
