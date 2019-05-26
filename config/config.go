package config

import "github.com/kelseyhightower/envconfig"

type environment struct {
	DB                string `required:"true"`
	S3AccessKeyID     string `required:"true"`
	S3SecretAccessKey string `required:"true"`
	S3Bucket          string `required:"true"`
	TokenSecretKey    string `required:"true"`
}

var _env *environment

func Env() *environment {
	return _env
}

func init() {
	var env environment
	if err := envconfig.Process("MIKANMUSIC", &env); err != nil {
		panic(err)
	}
	_env = &env
}
