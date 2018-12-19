package config

import (
	"github.com/hscasn/go-microservice/pkg/env"
)

// Framework is the top-level configuration struct
type Framework struct {
	Name string
}

// Create will recover the environment settings and parse them into a struct
func Create() Framework {
	return Framework{
		Name: env.String("FRAMEWORK_NAME") + "ok",
	}
}
