package service

import (
	"context"
	"os"
	"testing"

	"github.com/hserge/namak/model"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type EmailServiceTestSuite struct {
	suite.Suite
	dbPool *pgxpool.Pool
}

func TestEmailServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EmailServiceTestSuite))
}

func (s *EmailServiceTestSuite) SetupSuite() {
	var err error

	err = godotenv.Load("../.env")
	if err != nil {
		s.T().Log(err)
	}

	s.dbPool, err = pgxpool.New(context.TODO(), os.Getenv("TEST_DSN"))
	if err != nil {
		s.T().Log(err)
	}
}

func (s *EmailServiceTestSuite) TestEmailService_Create() {
	email := &model.Email{
		IsActive:  pgtype.Bool{Bool: true, Valid: true},
		Email:     pgtype.Text{String: "abc@email.com", Valid: true},
		FirstName: pgtype.Text{String: "Marco", Valid: true},
		LastName:  pgtype.Text{String: "Polo", Valid: true},
		Container: map[string]any{"a": "1", "b": "bee"},
	}
	svc := NewEmailService(context.TODO(), s.dbPool)
	err := svc.Create(email)
	s.NoError(err)
}

func (s *EmailServiceTestSuite) TestEmailService_Get() {
	svc := NewEmailService(context.TODO(), s.dbPool)
	_, err := svc.Get(1)
	s.NoError(err)
}
