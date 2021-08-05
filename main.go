package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
)

var (
	environment = os.Getenv("ENVIRONMENT")
	requiredEnvVars = []string{"ENVIRONMENT"}
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel) // Default log level is warn, unless debug flag is present
	for _, v := range requiredEnvVars {
		if _, ok := os.LookupEnv(v); !ok {
			log.Fatalf("environment variable %s is required", v)
		}
	}
}

func main() {
	lambda.Start(handler)
}
