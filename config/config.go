package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetRegion() string {
	return getEnvironmentValue("REGION")
}

func GetPort() int {
	portStr := getEnvironmentValue("PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}

func GetTables() []string {
	tablesStr := getEnvironmentValue("TABLES")
	return strings.Split(tablesStr, ",")
}
