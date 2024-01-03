package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

// noOpWriteSyncer is a WriteSyncer that does not perform the Sync operation.
type noOpWriteSyncer struct {
	zapcore.WriteSyncer
}

// Sync on noOpWriteSyncer is a no-op and always returns nil.
func (n noOpWriteSyncer) Sync() error {
	return nil
}

func InitLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	logFile, err := os.OpenFile("project_name.run.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// Wrap the os.Stdout with our noOpWriteSyncer
	consoleSyncer := noOpWriteSyncer{zapcore.AddSync(os.Stdout)}

	multiWriteSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(logFile),
		consoleSyncer,
	)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config.EncoderConfig),
		multiWriteSyncer,
		config.Level,
	)

	Logger = zap.New(core)
	// print logger init success
	Logger.Info("Booting up the application...")
}
