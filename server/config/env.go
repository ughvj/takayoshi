package config

import (
	"os"
	"flag"
)

type Env string

const (
	EnvTest Env = "test"
	EnvDev  Env = "dev"
	EnvStg  Env = "stg"
	EnvProd Env = "prod"
)

func CurrentEnv() Env {
	if IsTestEnv() {
		return EnvTest
	}
	env := os.Getenv("ENV")
	if env == "" {
		return EnvDev
	}
	return Env(env)
}

func IsTestEnv() bool {
	return flag.Lookup("test.v") != nil
}
