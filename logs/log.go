package logs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var config zap.Config

func init() {
	mode := os.Getenv("MODE")

	if mode == "Prod" {
		// NOTE : Config สำหรับ Production
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.StacktraceKey = "" // NOTE : ปิด Stack trace ได้เผื่อไม่อยากใช้
	} else {
		// NOTE : Config สำหรับ Development
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.ConsoleSeparator = " | "
	}

	var err error
	// NOTE : ถ้าไม่มี AddCallerSkip จะไม่ขึ้นเป็น บรรทัดที่เรียก log จริงๆ
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// NOTE : `message interface{}` เปรียบเหมือน Object แบบไม่เจาะจง type ใน Java
func Error(message interface{}, fields ...zap.Field) {
	// NOTE : Cast to error type or string
	switch v := message.(type) { // find type of `message`
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	default:
		log.Error("unkown error message type", fields...)
	}
}
