package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port string
	MongoUri string
	Env string
}

var Envs Env = Env{}

func (env *Env) Init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	env.Port = os.Getenv("PORT")
	env.MongoUri = os.Getenv("MONGO_URI")
	env.Env = os.Getenv("ENV")
}