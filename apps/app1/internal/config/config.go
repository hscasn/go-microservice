package config

import (
	"github.com/hscasn/go-microservice/pkg/env"
)

// Framework is the top-level configuration struct
type Framework struct {
	Name string
}

// New will recover the environment settings and parse them into a struct
func New() Framework {
	return Framework{
		Name: env.String("FRAMEWORK_NAME"),
	}
}
