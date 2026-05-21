package models

import "time"

type Scope struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	Target    string    `json:"target"`
	ScopeType string    `json:"scope_type"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
}
