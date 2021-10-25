package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	TimeInterval           string
	ApplicationPort        string
	FilePath 							 string
}

func ParseEnv(key string) string {
	_ = godotenv.Load()
	value := os.Getenv(key)
	if value == ""  {
		panic(fmt.Sprintf("Environment variable not found: %v", key))
	} 
	return value
}


//get param from .env
func GetEnvironment() *Environment {
	return &Environment{
		TimeInterval:        ParseEnv("TimeInterval"),
		ApplicationPort:     ParseEnv("ApplicationPort"),
		FilePath:            ParseEnv("FilePath"),
	}
}