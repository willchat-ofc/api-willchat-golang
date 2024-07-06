package helper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoContainer(t *testing.T) (*mongo.Database, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.4.6",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	host, err := mongoC.Host(ctx)
	require.NoError(t, err)

	port, err := mongoC.MappedPort(ctx, "27017")
	require.NoError(t, err)

	uri := "mongodb://" + host + ":" + port.Port()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	require.NoError(t, err)

	db := client.Database("testdb")

	return db, func() {
		client.Disconnect(ctx)
		mongoC.Terminate(ctx)
	}
}
