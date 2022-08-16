package logger

// 本文参考的是 Uber-go zap 日志库   https://www.liwenzhou.com/posts/Go/zap/
import "go.uber.org/zap"

var logger *zap.Logger

func initLogger() {
	logger, _ = zap.NewProduction()
}
