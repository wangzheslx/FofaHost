package log

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

func NewLogger() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", "123456",
		"service.name", "fofa",
		"service.version", "v0.0.1",
	)
}
