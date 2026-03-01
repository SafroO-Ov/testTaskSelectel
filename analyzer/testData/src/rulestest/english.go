package rulestest

import (
	"log/slog"

	"go.uber.org/zap"
)

// slog: глобальные функции

func testSlogGlobalEnglish() {
	slog.Debug("starting server") // ok
	slog.Info("starting server")  // ok
	slog.Warn("starting server")  // ok
	slog.Error("starting server") // ok

	slog.Debug("запуск сервера") // want "only English letters"
	slog.Info("запуск сервера")  // want "only English letters"
	slog.Warn("запуск сервера")  // want "only English letters"
	slog.Error("запуск сервера") // want "only English letters"
}

// slog: экземпляр Logger

func testSlogLoggerEnglish(logger *slog.Logger) {
	logger.Debug("starting server") // ok
	logger.Info("starting server")  // ok
	logger.Warn("starting server")  // ok
	logger.Error("starting server") // ok

	logger.Debug("запуск сервера") // want "only English letters"
	logger.Info("запуск сервера")  // want "only English letters"
	logger.Warn("запуск сервера")  // want "only English letters"
	logger.Error("запуск сервера") // want "only English letters"
}

// zap.Logger

func testZapLoggerEnglish(zlogger *zap.Logger) {
	zlogger.Debug("starting server") // ok
	zlogger.Info("starting server")  // ok
	zlogger.Warn("starting server")  // ok
	zlogger.Error("starting server") // ok

	zlogger.Debug("запуск сервера") // want "only English letters"
	zlogger.Info("запуск сервера")  // want "only English letters"
	zlogger.Warn("запуск сервера")  // want "only English letters"
	zlogger.Error("запуск сервера") // want "only English letters"
}

// zap.SugaredLogger: Debug/Info/Warn/Error

func testZapSugaredEnglish(slogger *zap.SugaredLogger) {
	// Debug/Info/Warn/Error (простые)
	slogger.Debug("starting server") // ok
	slogger.Info("starting server")  // ok
	slogger.Warn("starting server")  // ok
	slogger.Error("starting server") // ok

	slogger.Debug("запуск сервера") // want "only English letters"
	slogger.Info("запуск сервера")  // want "only English letters"
	slogger.Warn("запуск сервера")  // want "only English letters"
	slogger.Error("запуск сервера") // want "only English letters"

	// Debugf/Infof/Warnf/Errorf (форматные)
	slogger.Debugf("starting server on port %d", 8080) // ok
	slogger.Infof("starting server on port %d", 8080)  // ok
	slogger.Warnf("starting server on port %d", 8080)  // ok
	slogger.Errorf("starting server on port %d", 8080) // ok

	slogger.Debugf("запуск сервера на порту %d", 8080) // want "only English letters"
	slogger.Infof("запуск сервера на порту %d", 8080)  // want "only English letters"
	slogger.Warnf("запуск сервера на порту %d", 8080)  // want "only English letters"
	slogger.Errorf("запуск сервера на порту %d", 8080) // want "only English letters"

	// Debugw/Infow/Warnw/Errorw (with)
	slogger.Debugw("starting server", "port", 8080) // ok
	slogger.Infow("starting server", "port", 8080)  // ok
	slogger.Warnw("starting server", "port", 8080)  // ok
	slogger.Errorw("starting server", "port", 8080) // ok

	slogger.Debugw("запуск сервера", "port", 8080) // want "only English letters"
	slogger.Infow("запуск сервера", "port", 8080)  // want "only English letters"
	slogger.Warnw("запуск сервера", "port", 8080)  // want "only English letters"
	slogger.Errorw("запуск сервера", "port", 8080) // want "only English letters"
}
