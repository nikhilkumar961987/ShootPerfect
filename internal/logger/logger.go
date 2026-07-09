/******************************************************************************
 * Package logger provides application logging setup for ShootPerfect Core.
 *
 * The rest of the code should use this package instead of creating loggers
 * directly. This keeps logging consistent and makes it easier to change log
 * format or log level later.
 *****************************************************************************/
package logger

import (
	"log/slog"
	"os"
)

// New creates the default application logger.
//
// For now, it writes human-readable text logs to stdout at info level.
// Later this can be extended to support JSON logs, debug level, or config-based logging.
func New() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler)
}
