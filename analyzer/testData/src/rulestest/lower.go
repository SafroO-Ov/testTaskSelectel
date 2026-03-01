package rulestest

import (
	"log/slog"

	"go.uber.org/zap"
)

// ===== slog: глобальные функции =====

func testSlogGlobalLower() {
	// корректные — начинаются с маленькой буквы
	slog.Debug("starting server") // ok
	slog.Info("starting server")  // ok
	slog.Warn("starting server")  // ok
	slog.Error("starting server") // ok

	// некорректные — с заглавной
	slog.Debug("Starting server") // want "start with a lower-case letter"
	slog.Info("Starting server")  // want "start with a lower-case letter"
	slog.Warn("Starting server")  // want "start with a lower-case letter"
	slog.Error("Starting server") // want "start with a lower-case letter"
}

// ===== slog: экземпляр Logger =====

func testSlogLoggerLower(logger *slog.Logger) {
	logger.Debug("starting server") // ok
	logger.Info("starting server")  // ok
	logger.Warn("starting server")  // ok
	logger.Error("starting server") // ok

	logger.Debug("Starting server") // want "start with a lower-case letter"
	logger.Info("Starting server")  // want "start with a lower-case letter"
	logger.Warn("Starting server")  // want "start with a lower-case letter"
	logger.Error("Starting server") // want "start with a lower-case letter"
}

// ===== zap.Logger =====

func testZapLoggerLower(zlogger *zap.Logger) {
	zlogger.Debug("starting server", zap.String("env", "dev")) // ok
	zlogger.Info("starting server", zap.String("env", "dev"))  // ok
	zlogger.Warn("starting server", zap.String("env", "dev"))  // ok
	zlogger.Error("starting server", zap.String("env", "dev")) // ok

	zlogger.Debug("Starting server", zap.String("env", "prod")) // want "start with a lower-case letter"
	zlogger.Info("Starting server", zap.String("env", "prod"))  // want "start with a lower-case letter"
	zlogger.Warn("Starting server", zap.String("env", "prod"))  // want "start with a lower-case letter"
	zlogger.Error("Starting server", zap.String("env", "prod")) // want "start with a lower-case letter"
}

// ===== zap.SugaredLogger: Debug/Info/Warn/Error (+ f/w) =====

func testZapSugaredLower(slogger *zap.SugaredLogger) {
	// простые методы
	slogger.Debug("starting server") // ok
	slogger.Info("starting server")  // ok
	slogger.Warn("starting server")  // ok
	slogger.Error("starting server") // ok

	slogger.Debug("Starting server") // want "start with a lower-case letter"
	slogger.Info("Starting server")  // want "start with a lower-case letter"
	slogger.Warn("Starting server")  // want "start with a lower-case letter"
	slogger.Error("Starting server") // want "start with a lower-case letter"

	// форматные методы
	slogger.Debugf("starting server on port %d", 8080) // ok
	slogger.Infof("starting server on port %d", 8080)  // ok
	slogger.Warnf("starting server on port %d", 8080)  // ok
	slogger.Errorf("starting server on port %d", 8080) // ok

	slogger.Debugf("Starting server on port %d", 8080) // want "start with a lower-case letter"
	slogger.Infof("Starting server on port %d", 8080)  // want "start with a lower-case letter"
	slogger.Warnf("Starting server on port %d", 8080)  // want "start with a lower-case letter"
	slogger.Errorf("Starting server on port %d", 8080) // want "start with a lower-case letter"

	// w-методы
	slogger.Debugw("starting server", "port", 8080) // ok
	slogger.Infow("starting server", "port", 8080)  // ok
	slogger.Warnw("starting server", "port", 8080)  // ok
	slogger.Errorw("starting server", "port", 8080) // ok

	slogger.Debugw("Starting server", "port", 8080) // want "start with a lower-case letter"
	slogger.Infow("Starting server", "port", 8080)  // want "start with a lower-case letter"
	slogger.Warnw("Starting server", "port", 8080)  // want "start with a lower-case letter"
	slogger.Errorw("Starting server", "port", 8080) // want "start with a lower-case letter"
}
