package testhelper

import (
	"context"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupDockerNetwork(ctx context.Context) testcontainers.DockerNetwork {
	net, err := network.New(ctx, network.WithDriver("bridge"))

	if err != nil {
		log.Panic(err)
	}

	return *net
}

func setupDatabaseContainer(imageName string, port nat.Port, env map[string]string, ctx context.Context, network testcontainers.DockerNetwork, aliases []string) testcontainers.Container {

	containerReq := testcontainers.ContainerRequest{
		Image:        imageName,
		ExposedPorts: []string{port.Port() + "/" + port.Proto()},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections"),
			wait.ForListeningPort(port),
		),
		Networks: []string{network.Name},
		NetworkAliases: map[string][]string{
			network.Name: aliases,
		},
		Env: env,
	}

	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	if err != nil {
		log.Panic(err)
	}

	return container
}

func setupCustomContainer(filePath string, network testcontainers.DockerNetwork, aliases []string, env map[string]string, microPort nat.Port, dbPort nat.Port) testcontainers.Container {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{Context: filePath},
			Networks:       []string{network.Name},
			NetworkAliases: map[string][]string{network.Name: aliases},
			Env:            env,
			ExposedPorts:   []string{microPort.Port() + "/tcp"},
			WaitingFor:     wait.ForListeningPort(microPort),
		},
		Started: true,
	})

	if err != nil {
		log.Panic("error hit ", err)
	}

	return container
}

func SetupTestContainerEnvironment(ctx context.Context, hostIP string) string {

	dNetwork := setupDockerNetwork(ctx)

	postgresPort := nat.Port("5432/tcp")

	prescriptionDBContainer := setupDatabaseContainer("postgres", postgresPort, map[string]string{
		"POSTGRES_USER":     "postgres",
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_DB":       "prescription",
	}, ctx, dNetwork, []string{"db-rx-bridge"})

	authDBContainer := setupDatabaseContainer("postgres", postgresPort, map[string]string{
		"POSTGRES_USER":     "postgres",
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_DB":       "auth",
	}, ctx, dNetwork, []string{"db-auth-bridge"})

	prescriptionDBPort, err := prescriptionDBContainer.MappedPort(ctx, postgresPort)
	prescriptionDBHost, err := prescriptionDBContainer.ContainerIP(ctx)
	if err != nil {
		log.Panic("error trying to grab port for prescription db container", err)
	}

	authDBPort, err := authDBContainer.MappedPort(ctx, postgresPort)
	authDBHost, err := authDBContainer.ContainerIP(ctx)
	if err != nil {
		log.Panic("error trying to grab port for auth db container", err)
	}

	prescriptionContainer := setupCustomContainer("../../../../prescription", dNetwork, []string{"prescription-bridge"}, map[string]string{
		"POSTGRES_USER":     "postgres",
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_DB":       "prescription",
		"PORT":              "8080",
		"HOST":              prescriptionDBHost,
		"DB_PORT":           "5432",
	}, nat.Port("8080/tcp"), prescriptionDBPort)

	prescriptionMicroPort, err := prescriptionContainer.MappedPort(ctx, nat.Port("8080/tcp"))

	if err != nil {
		log.Panic("error trying to get prescription micro port ", err)
	}

	authContainer := setupCustomContainer("../../../../auth", dNetwork, []string{"auth-bridge"}, map[string]string{
		"POSTGRES_USER":     "postgres",
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_DB":       "auth",
		"PORT":              "8080",
		"HOST":              authDBHost,
		"JWT_SECRET":        "thisisajwtsecretbrod",
		"DB_PORT":           "5432",
	}, nat.Port("8080/tcp"), authDBPort)

	authMicroPort, err := authContainer.MappedPort(ctx, nat.Port("8080/tcp"))

	if err != nil {
		log.Panic("error trying to get auth micro port", err)
	}

	gatewayContainer := setupCustomContainer("../../../.", dNetwork, []string{""}, map[string]string{
		"PORT":               "8080",
		"PRESCRIPTION_MICRO": prescriptionMicroPort.Port(),
		"RX_HISTORY_MICRO":   "8006",
		"AUTH_MICRO":         authMicroPort.Port(),
		"JWT_SECRET":         "thisisajwtsecretbrod",
		"HOST_IP":            hostIP,
	}, nat.Port("8080/tcp"), nat.Port("5432/tcp"))

	gatewayPort, err := gatewayContainer.MappedPort(ctx, nat.Port("8080/tcp"))

	if err != nil {
		log.Panic("error trying to get gateway micro port", err)
	}
	return gatewayPort.Port()

}
