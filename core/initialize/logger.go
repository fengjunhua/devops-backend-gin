package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)
var (
	Logger *zap.Logger
)

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel)
	Logger = zap.New(core)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file,_ := os.Create("./devops.log")
	return zapcore.AddSync(file)
}