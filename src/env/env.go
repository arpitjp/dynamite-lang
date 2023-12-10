package env

import "os"

type rocketEnvVar string

const (
	ROCKET_TAB_WIDTH rocketEnvVar = "ROCKET_TAB_WIDTH"
)

var defaultValues = map[rocketEnvVar]string{
	ROCKET_TAB_WIDTH: "    ", // 1 tab = 4 spaces
}

func Getenv(key rocketEnvVar) string {
	val := os.Getenv(string(key))
	if val != "" {
		return val
	}
	return defaultValues[key]
}