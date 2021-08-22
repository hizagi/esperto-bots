package mongodb

import (
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
	image  = "mongo:4.4.3"
	logMsg = "database system is ready to accept connections"
)

//EmbeddedPostgres spins up a postgres container.
func EmbedMongo(t *testing.T, configuration viper.MongoConfiguration) (string, int) {
	t.Helper()
	ctx := context.Background()
	natPort := fmt.Sprintf("%d/tcp", 27017)
	// Setup and startup container
	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{natPort},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": configuration.Username,
			"MONGO_INITDB_ROOT_PASSWORD": configuration.Password,
			"MONGO_INITDB_DATABASE":      configuration.DatabaseName,
		},
		WaitingFor: wait.ForLog(logMsg).
			WithPollInterval(100 * time.Millisecond).
			WithOccurrence(2),
	}
	mongo, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		t.Error(err)
	}
	// Even after log message found Postgres needs a touch more...
	time.Sleep(200 * time.Millisecond)
	// When test is done terminate container
	t.Cleanup(func() {
		_ = mongo.Terminate(ctx)
	})
	// Get the container info needed
	containerPort, err := mongo.MappedPort(ctx, nat.Port(natPort))
	if err != nil {
		t.Error(err)
	}
	containerHost, err := mongo.Host(ctx)
	if err != nil {
		t.Error(err)
	}

	return containerHost, containerPort.Int()
}
