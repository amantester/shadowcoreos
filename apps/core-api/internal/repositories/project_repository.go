package repositories

import (
	"context"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	DB *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	query := `
	INSERT INTO projects (name, description, created_by)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`

	return r.DB.QueryRow(
		context.Background(),
		query,
		project.Name,
		project.Description,
		project.CreatedBy,
	).Scan(&project.ID, &project.CreatedAt)
}

func (r *ProjectRepository) List() ([]models.Project, error) {
	query := `
	SELECT id, name, description, created_by, created_at
	FROM projects
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		var p models.Project

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.CreatedBy,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		projects = append(projects, p)
	}

	return projects, nil
}
