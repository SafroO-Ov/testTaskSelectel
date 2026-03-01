// testdata/src/rulestest/sensitive.go
package rulestest

import (
	"log/slog"

	"go.uber.org/zap"
)

// ===== slog: глобальные функции =====

func testSlogGlobalSensitive(password, token, apiKey string) {
	// позитивные кейсы
	slog.Debug("user authenticated successfully") // ok
	slog.Info("api request completed")            // ok
	slog.Warn("token validated")                  // ok
	slog.Error("user session created")            // ok

	// негативные кейсы
	slog.Debug("user password: " + password)              // want "sensitive data"
	slog.Info("api_key=" + apiKey)                        // want "sensitive data"
	slog.Warn("token: " + token)                          // want "sensitive data"
	slog.Error("user password is incorrect: " + password) // want "sensitive data"
}

// ===== slog: экземпляр Logger =====

func testSlogLoggerSensitive(logger *slog.Logger, password, token, apiKey string) {
	// позитивные
	logger.Debug("user authenticated successfully") // ok
	logger.Info("api request completed")            // ok
	logger.Warn("token validated")                  // ok
	logger.Error("user session created")            // ok

	// негативные
	logger.Debug("user password: " + password)              // want "sensitive data"
	logger.Info("api_key=" + apiKey)                        // want "sensitive data"
	logger.Warn("token: " + token)                          // want "sensitive data"
	logger.Error("user password is incorrect: " + password) // want "sensitive data"
}

// ===== zap.Logger =====

func testZapLoggerSensitive(zlogger *zap.Logger, password, token, apiKey string) {
	// позитивные
	zlogger.Debug("user authenticated successfully", zap.String("user", "john")) // ok
	zlogger.Info("api request completed", zap.String("path", "/login"))          // ok
	zlogger.Warn("token validated", zap.String("status", "ok"))                  // ok
	zlogger.Error("user session created", zap.String("user", "john"))            // ok

	// негативные
	zlogger.Debug("user password: "+password, zap.String("user", "john"))              // want "sensitive data"
	zlogger.Info("api_key="+apiKey, zap.String("service", "payments"))                 // want "sensitive data"
	zlogger.Warn("token: "+token, zap.String("user", "john"))                          // want "sensitive data"
	zlogger.Error("user password is incorrect: "+password, zap.String("user", "john")) // want "sensitive data"
}

// ===== zap.SugaredLogger =====

func testZapSugaredSensitive(slogger *zap.SugaredLogger, password, token, apiKey string) {
	// позитивные
	slogger.Debug("user authenticated successfully") // ok
	slogger.Info("api request completed")            // ok
	slogger.Warn("token validated")                  // ok
	slogger.Error("user session created")            // ok

	// негативные
	slogger.Debug("user password: " + password)              // want "sensitive data"
	slogger.Info("api_key=" + apiKey)                        // want "sensitive data"
	slogger.Warn("token: " + token)                          // want "sensitive data"
	slogger.Error("user password is incorrect: " + password) // want "sensitive data"
}
