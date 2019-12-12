package conf

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/_metalogic_/log"
)

const secretsDir = "/var/run/secrets"

// IfGetenv returns the value of environment variable name if found, else deflt
func IfGetenv(name, deflt string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		value = deflt
	}
	return value
}

// MustGetenv returns the value of environment variable name;
// if name is not found exit with fatal error
func MustGetenv(name string) (v string) {
	v = os.Getenv(name)
	if v == "" {
		log.Fatalf("environment variable '%s' is not set", name)
	}
	return v
}

// MustGetInt returns the int value of environment variable name;
// if name is not found or value cannot be parsed as an int exit with fatal error
func MustGetInt(name string, deflt int) (value int) {
	env := os.Getenv(name)
	if env != "" {
		var err error
		value, err = strconv.Atoi(env)
		if err != nil {
			log.Fatal(err.Error())
		}
		return value
	}
	return deflt
}

// MustGetDuration returns the time.Duration value of environment variable name;
// if name is not found or value cannot be parsed as a time.Duration exit with fatal error
func MustGetDuration(name string, deflt time.Duration) (value time.Duration) {
	env := os.Getenv(name)
	if env != "" {
		var err error
		value, err = time.ParseDuration(env)
		if err != nil {
			log.Fatal(err.Error())
		}
		return value
	}
	return deflt
}

// GetSecret returns the value of Docker secret name
func GetSecret(name string) (secret string, err error) {
	name = os.Getenv("ENV") + "_" + name
	name = strings.ToUpper(name)
	log.Info("GetSecret has " + name)
	bytes, err := ioutil.ReadFile(secretsDir + name)
	if err != nil {
		return secret, err
	}
	log.Info("string from GetSecret is " + string(bytes))
	return string(bytes), nil
}

// MustGetSecret returns the value of the Docker secret named  $ENV_name;
// if the environment variable ENV is empty or if the secret cannot be read exit with fatal error
func MustGetSecret(name string) string {
	secret, err := GetSecret(name)
	if err != nil || secret == "" {
		log.Fatalf("secret " + name + " not configured")
	}
	return secret
}

// MustGetConfig returns the value of the Docker secret named $ENV_name, if it exists;
// if the environment variable ENV is empty or if the secret cannot be read, returns
// the value of environment variable name; if name is not found in either Docker secrets
// or as an environment variable, exit with fatal error
func MustGetConfig(name string) string {
	config, err := GetSecret(name)
	log.Debug("%v", config)
	if err != nil || config == "" {
		config := os.Getenv(name)
		if config == "" {
			log.Fatalf(name + " not configured")
		}
	}
	log.Debug("%v", name)
	return config
}
