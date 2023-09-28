package utilities

import (
	"github.com/joho/godotenv"
	//"os"
)
//this is not used till now
type GetEnv struct{}

var env map[string]string

func LocalEnvInitialize() {
	// env_selection := os.Getenv("APP_ENV")
	// env_selection := "local"
	//envPath := "./environments/" + env_selection + ".env"
	envPath := "./environments/" + ".env"
	environment, err := godotenv.Read(envPath)
	if err != nil {

	}
	env = environment
}
func ENV() map[string]string {
	return env
}
