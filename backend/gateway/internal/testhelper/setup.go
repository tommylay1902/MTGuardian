package testhelper

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupDockerNetwork() testcontainers.DockerNetwork {
	ctx := context.Background()

	net, err := network.New(ctx, network.WithDriver("bridge"))

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := net.Remove(ctx); err != nil {
			log.Panic(err)
		}
	}()

	return *net
}

func SetupTestingContainer(imageName string, port nat.Port, env map[string]string, dNetwork testcontainers.DockerNetwork) testcontainers.Container {
	fmt.Println("HELLOO", port.Port(), port.Proto())
	// 1. init gateway
	containerGatewayReq := testcontainers.ContainerRequest{
		Image:        imageName,
		ExposedPorts: []string{port.Port() + "/" + port.Proto()},
		WaitingFor:   wait.ForListeningPort(port),
		Networks:     []string{dNetwork.Name},
		Env:          env,
	}

	dbGatewayContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerGatewayReq,
			Started:          true,
		})

	if err != nil {
		log.Panic(err)
	}

	return dbGatewayContainer
}
