package duit_utils

import "os"

var (
	RuntimeValidationEnabled = true
)

func CheckEnv() {
	val, ok := os.LookupEnv("RUNTIME_VALIDATION_ENABLED")
	if !ok {
		return
	}

	RuntimeValidationEnabled = val == "true" || val == "1"
}