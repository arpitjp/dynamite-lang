package parser

import (
	"dynamite/src/env"
	"dynamite/src/logger"
	"fmt"
	"runtime"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace() string {
	if tracingEnabled := env.Getenv(env.DYNAMITE_PARSER_TRACING_ENABLED); tracingEnabled == "" {
		return ""
	}
	fxnName := getFuncName()
	incIdent()
	tracePrint(logger.Info("BEGIN ") + fxnName)
	return fxnName
}

func untrace(fxnName string) {
	if tracingEnabled := env.Getenv(env.DYNAMITE_PARSER_TRACING_ENABLED); tracingEnabled == "" {
		return
	}
	tracePrint(logger.Error("END ") + fxnName)
	decIdent()
}

func getFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fullName := details.Name()
		arr := strings.Split(fullName, ".")
		return arr[len(arr) - 1]
	}
	return "FAILED IN GETTING FUNCTION NAME"
}