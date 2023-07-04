package main

import (
	"os"
	"strconv"
)

func getEnv(name string, def string) string {
	value := os.Getenv(name)
	if value == "" {
		return def
	}

	return value
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
