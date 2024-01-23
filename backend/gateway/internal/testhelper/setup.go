package testhelper

// func SetupDockerNetwork(ctx context.Context) testcontainers.DockerNetwork {

// 	net, err := network.New(ctx, network.WithDriver("bridge"))

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return *net
// }

// func SetupTestingContainer(imageName string, port nat.Port, env map[string]string, dNetwork testcontainers.DockerNetwork, aliases []string, ctx context.Context) testcontainers.Container {

// 	containerReq := testcontainers.ContainerRequest{
// 		Image:        imageName,
// 		ExposedPorts: []string{port.Port() + "/" + port.Proto()},
// 		WaitingFor:   wait.ForListeningPort(port),
// 		Networks:     []string{dNetwork.Name},
// 		NetworkAliases: map[string][]string{
// 			dNetwork.Name: aliases,
// 		},
// 		Env: env,
// 	}

// 	dbGatewayContainer, err := testcontainers.GenericContainer(
// 		ctx,
// 		testcontainers.GenericContainerRequest{
// 			ContainerRequest: containerReq,
// 			Started:          true,
// 		})

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return dbGatewayContainer
// }

import (
	"context"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// func SetupDockerNetwork(ctx context.Context) testcontainers.DockerNetwork {
// 	net, err := network.New(ctx, network.WithDriver("driver"))

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return *net
// }

func SetupTestingContainer(imageName string, port nat.Port, env map[string]string, ctx context.Context) testcontainers.Container {

	containerReq := testcontainers.ContainerRequest{
		Image:        imageName,
		ExposedPorts: []string{port.Port() + "/" + port.Proto()},
		WaitingFor:   wait.ForListeningPort(port),
		// Networks:     []string{dNetwork.Name},
		// NetworkAliases: map[string][]string{
		// 	dNetwork.Name: aliases,
		// },
		Env: env,
	}

	dbGatewayContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	if err != nil {
		log.Panic(err)
	}

	return dbGatewayContainer
}
