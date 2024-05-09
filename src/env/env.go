package env

import "os"

type rocketEnvVar string

const (
	DYNAMITE_TAB_WIDTH rocketEnvVar = "DYNAMITE_TAB_WIDTH"
	DYNAMITE_PARSER_TRACING_ENABLED rocketEnvVar = "DYNAMITE_PARSER_TRACING"
	DYNAMITE_LEXER_INSPECT_TOKESN rocketEnvVar = "DYNAMITE_LEXER_INSPECT_TOKESN"
)

var defaultValues = map[rocketEnvVar]string{
	DYNAMITE_TAB_WIDTH: "    ", // 1 tab = 4 spaces
	DYNAMITE_PARSER_TRACING_ENABLED: "",
	DYNAMITE_LEXER_INSPECT_TOKESN: "d",
}

func Getenv(key rocketEnvVar) string {
	val := os.Getenv(string(key))
	if val != "" {
		return val
	}
	return defaultValues[key]
}