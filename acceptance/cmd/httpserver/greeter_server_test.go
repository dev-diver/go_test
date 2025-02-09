package main_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"valentine.com/tdd/acceptance/adapters/httpserver"
	specifications "valentine.com/tdd/acceptance/specification"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../../.",
			Dockerfile: "./acceptance/cmd/httpserver/Dockerfile",
			// set to false if you want less spam, but this is helpful if you're having troubles
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	driver := httpserver.Driver{BaseURL: "http://localhost:8080", Client: &client}
	specifications.GreetSpecification(t, driver)
}
