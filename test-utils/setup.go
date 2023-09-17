package testutils

import (
	"context"
	"fmt"

	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

const (
	DbName     = ""
	DbUser     = ""
	DbPassword = ""
)

var Pport int

var DbContext *gorm.DB
var Container testcontainers.Container

func CreateContainer(ctx context.Context) (testcontainers.Container, error) {
	fmt.Println("creating db container")

	var env = map[string]string{
		"POSTGRES_PASSWORD": DbPassword,
		"POSTGRES_USER":     DbUser,
		"POSTGRES_DB":       DbName,
	}
	var port = "5432/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:14-alpine",
			ExposedPorts: []string{port},
			Env:          env,
			WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "5432")
	Pport = p.Int()
	if err != nil {
		return container, fmt.Errorf("failed to get container external port: %v", err)
	}

	fmt.Println("postgres container ready and running at port: ", p.Port())

	time.Sleep(time.Second)

	Container = container

	return container, nil
}

func TearDown() error {
	fmt.Println("tearing down db container")
	err := Container.Terminate(context.Background())
	return err
}
