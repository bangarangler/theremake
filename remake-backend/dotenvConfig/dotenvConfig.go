package dotenvConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/bangarangler/theremake/remake-backend/internal/projectpath"
	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	envPath := fmt.Sprintf("%s/.env", projectpath.Root)
	err := godotenv.Load(envPath)
	// err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

// var PgConnStr = goDotEnvVar("POSTGRES_URL")
var (
	Port      = goDotEnvVar("PORT")
	EmailUser = goDotEnvVar("EMAIL_USER")
	EmailPW   = goDotEnvVar("EMAIL_PW")
)
