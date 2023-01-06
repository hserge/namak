package service

import (
	"github.com/hserge/namak/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmailService struct {
	dbPool *pgxpool.Pool
}

func NewEmailService(dbPool *pgxpool.Pool) *EmailService {
	return &EmailService{dbPool: dbPool}
}

func (es *EmailService) Create(email *model.Email) error {
	query := `
        INSERT INTO emails (id, name, occupation) 
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at`
}
