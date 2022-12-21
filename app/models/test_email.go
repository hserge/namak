package models

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type EmailModelSuite struct {
	suite.Suite
	db *pgx.Conn
}

func TestEmailSuite(t *testing.T) {
	suite.Run(t, new(EmailModelSuite))
}

func (s *EmailModelSuite) SetupSuite() {
	var err error
	s.db, err = pgx.Connect(context.TODO(), dsn)
}
