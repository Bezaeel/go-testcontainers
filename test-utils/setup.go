package testutils

import (
	"context"
	"fmt"

	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbName     = "db123"
	dbUser     = "user123"
	dbPassword = "password123"
)

var container *postgres.PostgresContainer

func CreatePgSQLContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	fmt.Println("creating db container")

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create container: %v", err)
	}

	time.Sleep(time.Second)

	container = postgresContainer

	return postgresContainer, nil
}

func TearDown() error {
	fmt.Println("tearing down db container")
	err := container.Terminate(context.Background())
	return err
}
