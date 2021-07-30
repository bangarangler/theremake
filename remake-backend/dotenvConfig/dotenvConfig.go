package dotenvConfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load("./remake-backend/.env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

// var PgConnStr = goDotEnvVar("POSTGRES_URL")
var (
	Port = goDotEnvVar("PORT")
)
