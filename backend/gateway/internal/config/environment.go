package config

import (
	"log"
	"os"
)

func Setup() (string, string, string, string, string, string) {

	//setup jwt helper
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("jwt_secret is not specified")
	}

	//setup and return port
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not specified")
	}

	//setup and return port
	hostIP := os.Getenv("HOST_IP")

	if hostIP == "" {
		log.Fatal("host_ip is not specified")
	}

	pMicro := os.Getenv("PRESCRIPTION_MICRO")

	if pMicro == "" {
		log.Fatal("prescription_micro is not specified")
	}

	hMicro := os.Getenv("RX_HISTORY_MICRO")

	if hMicro == "" {
		log.Fatal("rx_history_micro is not specified")
	}

	aMicro := os.Getenv("AUTH_MICRO")

	if aMicro == "" {
		log.Fatal("auth_micro is not specified")
	}

	return secret, portString, hostIP, pMicro, hMicro, aMicro
}
