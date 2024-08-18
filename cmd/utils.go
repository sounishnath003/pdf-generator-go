package main

import (
	"log"
	"os"
)

type EnvOpts struct {
	PublicHost string
	Port       string
}

func InitEnvs() *EnvOpts {
	return &EnvOpts{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnv("APP_PORT", "3000"),
	}
}

func getEnv(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Printf("key=%s is not found. setting up fallback.value=%s\n", key, fallback)
		return fallback
	}
	return value
}
