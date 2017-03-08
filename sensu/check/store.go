package check

import stdCheck "github.com/gbolo/go-sensu-tls/sensu/check"

var Store = make(map[string]Check)

type Check interface {
	Execute() stdCheck.CheckOutput
}
