package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port string
}

var Envs Env = Env{}

func (env *Env) Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	env.Port = os.Getenv("PORT")
}