package handler

import (
	"fmt"

	stdCheck "github.com/gbolo/go-sensu-tls/sensu/check"
	"github.com/gbolo/go-sensu-client-tls/sensu/check"
)

func Ok(message string) check.ExtensionCheckResult {
	return check.ExtensionCheckResult{
		Status: stdCheck.Success,
		Output: fmt.Sprintf("OK: %s", message),
	}
}

func Warning(message string) check.ExtensionCheckResult {
	return check.ExtensionCheckResult{
		Status: stdCheck.Warning,
		Output: fmt.Sprintf("WARNING: %s", message),
	}
}

func Error(message string) check.ExtensionCheckResult {
	return check.ExtensionCheckResult{
		Status: stdCheck.Error,
		Output: fmt.Sprintf("ERROR: %s", message),
	}
}
