package rulestest

import (
	"log/slog"

	"go.uber.org/zap"
)

// ===== slog: глобальные функции =====

func testSlogGlobalSpecial() {
	// корректные: без спецсимволов/эмодзи
	slog.Debug("server started")                // ok
	slog.Info("connection failed")              // ok
	slog.Warn("something went wrong")           // ok
	slog.Error("failed to connect to database") // ok

	// некорректные: спецсимволы и эмодзи
	slog.Debug("server started!")                 // want "special characters or emojis"
	slog.Info("connection failed!!!")             // want "special characters or emojis"
	slog.Warn("warning: something went wrong...") // want "special characters or emojis"
	slog.Error("failed to connect to database 🚀") // want "special characters or emojis"
}

// ===== slog: экземпляр Logger =====

func testSlogLoggerSpecial(logger *slog.Logger) {
	logger.Debug("server started")                // ok
	logger.Info("connection failed")              // ok
	logger.Warn("something went wrong")           // ok
	logger.Error("failed to connect to database") // ok

	logger.Debug("server started!")                 // want "special characters or emojis"
	logger.Info("connection failed!!!")             // want "special characters or emojis"
	logger.Warn("warning: something went wrong...") // want "special characters or emojis"
	logger.Error("failed to connect to database 😱") // want "special characters or emojis"
}

// ===== zap.Logger =====

func testZapLoggerSpecial(zlogger *zap.Logger) {
	zlogger.Debug("server started", zap.String("env", "dev"))                // ok
	zlogger.Info("connection failed", zap.Int("code", 500))                  // ok
	zlogger.Warn("something went wrong", zap.String("component", "db"))      // ok
	zlogger.Error("failed to connect to database", zap.String("db", "main")) // ok

	zlogger.Debug("server started!", zap.String("env", "prod"))                // want "special characters or emojis"
	zlogger.Info("connection failed!!!", zap.Int("code", 500))                 // want "special characters or emojis"
	zlogger.Warn("warning: something went wrong...", zap.String("c", "db"))    // want "special characters or emojis"
	zlogger.Error("failed to connect to database 🚀", zap.String("db", "main")) // want "special characters or emojis"
}

// ===== zap.SugaredLogger: Debug/Info/Warn/Error (+ f/w) =====

func testZapSugaredSpecial(slogger *zap.SugaredLogger) {
	// простые методы
	slogger.Debug("server started")                // ok
	slogger.Info("connection failed")              // ok
	slogger.Warn("something went wrong")           // ok
	slogger.Error("failed to connect to database") // ok

	slogger.Debug("server started!")                 // want "special characters or emojis"
	slogger.Info("connection failed!!!")             // want "special characters or emojis"
	slogger.Warn("warning: something went wrong...") // want "special characters or emojis"
	slogger.Error("failed to connect to database 🤯") // want "special characters or emojis"

	// форматные методы
	slogger.Debugf("server started on port %d", 8080) // ok
	slogger.Infof("server started on port %d", 8080)  // ok
	slogger.Warnf("server started on port %d", 8080)  // ok
	slogger.Errorf("server started on port %d", 8080) // ok

	slogger.Debugf("server started on port %d!!!", 8080) // want "special characters or emojis"
	slogger.Infof("server started on port %d 🚀", 8080)   // want "special characters or emojis"
	slogger.Warnf("warning... server on port %d", 8080)  // want "special characters or emojis"
	slogger.Errorf("server on port %d!!!", 8080)         // want "special characters or emojis"

	// w-методы
	slogger.Debugw("server started", "port", 8080)                // ok
	slogger.Infow("connection failed", "port", 8080)              // ok
	slogger.Warnw("something went wrong", "port", 8080)           // ok
	slogger.Errorw("failed to connect to database", "port", 8080) // ok

	slogger.Debugw("server started!", "port", 8080)                 // want "special characters or emojis"
	slogger.Infow("connection failed!!!", "port", 8080)             // want "special characters or emojis"
	slogger.Warnw("warning: something went wrong...", "port", 8080) // want "special characters or emojis"
	slogger.Errorw("failed to connect to database 💥", "port", 8080) // want "special characters or emojis"
}
