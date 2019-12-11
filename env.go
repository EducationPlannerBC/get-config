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

func ifGetenv(name, deflt string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		value = deflt
	}
	return value
}

func mustGetenv(name string) (v string) {
	v = os.Getenv(name)
	if v == "" {
		log.Fatalf("environment variable '%s' is not set", name)
	}
	return v
}

func mustGetInt(name string, deflt int) (value int) {
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

func mustGetDuration(name string, deflt time.Duration) (value time.Duration) {
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

func getSecret(name string) (secret string, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadFile("/var/run/secrets/" + name)
	if err != nil {
		return secret, err
	}
	return string(bytes), nil
}

func mustGetSecret(name string) string {
	name = mustGetenv("ENV") + "_" + name
	name = strings.ToUpper(name)
	bytes, err := ioutil.ReadFile(secretsDir + name)
	if err != nil {
		log.Fatalf("secret " + name + " not configured")
	}
	return string(bytes)
}
