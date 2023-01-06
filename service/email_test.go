package service

import (
	"context"
	"os"
	"testing"

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
	s.dbPool, err = pgxpool.New(context.TODO(), os.Getenv("TEST_PGSQL_DSN"))
	if err != nil {
		panic(err)
	}
}

func (s *EmailServiceTestSuite) TestEmailService_Get() {
	svc := NewEmailService(context.TODO(), s.dbPool)
	_, err := svc.Get(1)
	s.NoError(err)
}
