package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "Level",
	})
	f, _ := os.Create("zap-12-05.log")
	enab := zapcore.ErrorLevel
	core := zapcore.NewCore(enc, f, enab)

	l := zap.New(core)

	l.Debug("debug", zap.Int8("key", 1))
	l.Panic("panic ", zap.String("key", "value"))

}
