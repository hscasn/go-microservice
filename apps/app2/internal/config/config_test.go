package config

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	// Setting env variables
	os.Setenv("FRAMEWORK_NAME", "hello_there")

	c := Create()
	if c.Name != "hello_there" {
		t.Errorf("framework name is incorrect")
	}
}
