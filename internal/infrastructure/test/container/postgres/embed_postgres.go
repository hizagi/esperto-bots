package postgres

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	image  = "postgres:12.4-alpine"
	logMsg = "database system is ready to accept connections"
)

//EmbedPostgres spins up a postgres container.
func EmbedPostgres(t *testing.T, configuration viper.PostgresConfiguration) (string, int) {
	t.Helper()

	ctx := context.Background()
	natPort := fmt.Sprintf("%d/tcp", 5432)
	// Setup and startup container
	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{natPort},
		Env: map[string]string{
			"POSTGRES_USER":     configuration.Username,
			"POSTGRES_PASSWORD": configuration.Password,
			"POSTGRES_DATABASE": configuration.DatabaseName,
		},
		WaitingFor: wait.ForLog(logMsg).
			WithPollInterval(100 * time.Millisecond).
			WithOccurrence(2),
	}
	pg, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		reader, _ := pg.Logs(ctx)
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)

		fmt.Printf("CONTAINER LOGS: %s", buf.String())

		t.Error(err)
	}
	// Even after log message found Postgres needs a touch more...
	time.Sleep(200 * time.Millisecond)
	// When test is done terminate container
	t.Cleanup(func() {
		_ = pg.Terminate(ctx)
	})
	// Get the container info needed
	containerPort, err := pg.MappedPort(ctx, nat.Port(natPort))
	if err != nil {
		t.Error(err)
	}
	containerHost, err := pg.Host(ctx)
	if err != nil {
		t.Error(err)
	}

	return containerHost, containerPort.Int()
}
