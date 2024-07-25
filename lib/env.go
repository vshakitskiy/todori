package lib

import (
	"log"
	"os"
)

func Getenv(key string) string {
	value, isSet := os.LookupEnv(key)

	if !isSet {
		log.Fatal("Environment variable not set: " + key)
	}

	return value
}