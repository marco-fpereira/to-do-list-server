package logger

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
)

var log *zap.Logger

func InitLogger() {
	logConfiguration := zap.Config{
		OutputPaths: []string{getOutputPath()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfiguration.Build()
}

func Info(ctx context.Context, message string, tags ...zap.Field) {
	tags = getDefaultTagsFromContext(ctx, tags...)
	log.Info(message, tags...)
	log.Sync()
}

func Error(ctx context.Context, message string, err error, tags ...zap.Field) {
	tags = getDefaultTagsFromContext(ctx, tags...)
	log.Error(message, append(tags, zap.NamedError("error", err))...)
	log.Sync()
}

func Fatal(ctx context.Context, message string, err error, tags ...zap.Field) {
	tags = getDefaultTagsFromContext(ctx, tags...)
	log.Fatal(message, append(tags, zap.NamedError("error", err))...)
	log.Sync()
}

func GetErrorFields(fields []string) []zap.Field {
	var zapFields []zap.Field
	for _, field := range fields {
		zapFields = append(zapFields, zap.String("field", field))
	}
	return zapFields
}

func getDefaultTagsFromContext(ctx context.Context, tags ...zap.Field) []zap.Field {
	if md, exists := metadata.FromOutgoingContext(ctx); exists {
		tags = append(tags, zap.String("requestId", strings.Join(md.Get("requestId"), "")))
	}
	return tags
}

func getOutputPath() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}
