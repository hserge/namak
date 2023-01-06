package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ContainerTestSuite struct {
	suite.Suite
	MyVar int
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ContainerTestSuite))
}

func (s *ContainerTestSuite) SetupTest() {
	s.MyVar = 5
	err := godotenv.Load("../.env")
	if err != nil {
		s.T().Fatal("Error loading .env file")
	}
}

func (s *ContainerTestSuite) TestNew() {
	ctx, log, pool, err := GetDefaults(GetDefaultParams{})
	s.NoError(err)

	container := New(ctx, log, pool)
	s.IsType(context.Background(), container.Ctx)
	s.Greater(int8(0), container.Log.GetLevel())
	s.NoError(container.Db.Ping(container.Ctx))
}

func (s *ContainerTestSuite) TestNewError() {
	s.Panics(func() { GetDefaults(GetDefaultParams{dsn: "invalid dsn"}) })
}
