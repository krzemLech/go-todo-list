package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	Port string
	MongoUri string
	Env string
	ProfaneWords string
	MaxTodos int
}

var Envs Env = Env{}

func (env *Env) Init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	var err error
	env.Port = os.Getenv("PORT")
	env.MongoUri = os.Getenv("MONGO_URI")
	env.Env = os.Getenv("ENV")
	env.ProfaneWords = os.Getenv("PROFANE_WORDS")
	env.MaxTodos, err = strconv.Atoi(os.Getenv("MAX_TODOS"))
	if err != nil {
		env.MaxTodos = 10
	}
}