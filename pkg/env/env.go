package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StringWithDefault gets an environment variable, giving a default if empty
func StringWithDefault(name string, def string) string {
	v := os.Getenv(name)
	if len(v) < 1 {
		return def
	}
	return v
}

// String gets an environment variable, panicking if it is empty
func String(name string) string {
	v := StringWithDefault(name, "")
	if len(v) < 1 {
		panic(fmt.Sprintf("Env variable '%s' is empty", name))
	}
	return v
}

// IntWithDefault gets an integer env variable, giving a default if empty
func IntWithDefault(name string, def int64) int64 {
	vs := StringWithDefault(name, "")
	if len(vs) < 1 {
		return def
	}
	v, err := strconv.ParseInt(vs, 10, 64)
	if err != nil {
		return def
	}
	return v
}

// Int gets an integer env variable, panicking if it is empty
func Int(name string) int64 {
	vs := StringWithDefault(name, "")
	if len(vs) < 1 {
		panic(fmt.Sprintf("Env variable '%s' is empty", name))
	}
	v, err := strconv.ParseInt(vs, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Env variable '%s' is not an integer", name))
	}
	return v
}

// FloatWithDefault gets a float env variable, giving a default if empty
func FloatWithDefault(name string, def float64) float64 {
	vs := StringWithDefault(name, "")
	if len(vs) < 1 {
		return def
	}
	v, err := strconv.ParseFloat(vs, 64)
	if err != nil {
		return def
	}
	return v
}

// Float gets a float env variable, panicking if it is empty
func Float(name string) float64 {
	vs := StringWithDefault(name, "")
	if len(vs) < 1 {
		panic(fmt.Sprintf("Env variable '%s' is empty", name))
	}
	v, err := strconv.ParseFloat(vs, 64)
	if err != nil {
		panic(fmt.Sprintf("Env variable '%s' is not a float", name))
	}
	return v
}

// BoolWithDefault gets a bool env variable, giving a default if empty
func BoolWithDefault(name string, def bool) bool {
	v := StringWithDefault(name, "")
	if len(v) < 1 {
		return def
	}
	return strings.ToLower(strings.Trim(v, " ")) == "true"
}

// Bool gets a bool env variable, panicking if it is empty
func Bool(name string) bool {
	v := StringWithDefault(name, "")
	if len(v) < 1 {
		panic(fmt.Sprintf("Env variable '%s' is empty", name))
	}
	return strings.ToLower(strings.Trim(v, " ")) == "true"
}
