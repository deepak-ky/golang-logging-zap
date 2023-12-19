package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var applicationLogger *zap.Logger

func init() {
	config := zap.Config{
		// Level is the minimum enabled logging level
		Level: zap.NewAtomicLevelAt(zap.DebugLevel),
		// Development puts the logger in development mode, which changes the
		// behavior of DPanicLevel and takes stacktraces more liberally.
		Development: false,
		// DisableCaller stops annotating logs with the calling function's file
		// name and line number. By default, all logs are annotated.
		DisableCaller: true,
		// DisableStacktrace completely disables automatic stacktrace capturing. By
		// default, stacktraces are captured for WarnLevel and above logs in
		// development and ErrorLevel and above in production.
		DisableStacktrace: true,
		// Sampling sets a sampling policy. A nil SamplingConfig disables sampling.
		// Zap allows you to control the volume of log entries by implementing sampling.
		// Sampling is the process of deciding which log statements to record and which ones to skip.
		// This can be useful in high-traffic applications where logging every statement could lead to a
		// significant performance overhead.
		// Sampling nil implies that every log statement will be recorded without any filtering based on a sampling policy.
		Sampling: nil,
		// Encoding sets the logger's encoding. Valid values are "json" and
		// "console
		Encoding: "json",
		// EncoderConfig sets options for the chosen encoder
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "timestamp",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			LevelKey:      "level",
			EncodeLevel:   zapcore.CapitalLevelEncoder, // EncodeLevel: zapcore.CapitalColorLevelEncoder with Encoding: "console" to print color value,
			MessageKey:    "message",
			StacktraceKey: "stacktrace",               // this will not work because DisableStacktrace: true
			CallerKey:     "invoke-point",             // this will not work because DisableCaller: true
			EncodeCaller:  zapcore.ShortCallerEncoder, // this will not work because DisableCaller: true
		},
		// OutputPaths is a list of URLs or file paths to write logging output to.
		OutputPaths: []string{"stdout", "logging.log"},
		// ErrorOutputPaths is a list of URLs to write internal logger errors to.
		ErrorOutputPaths: []string{"stderr"},
		// InitialFields is a collection of fields to add to the root logger.
		/* InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		}, */
	}

	applicationLogger = zap.Must(config.Build())
}

func main() {
	applicationLogger.Debug("this is DEBUG message", zap.String("url", "www.deepak.com"))
	applicationLogger.Info("this is INFO message", zap.String("url", "www.deepak.com"))
	applicationLogger.Warn("this is WARNING message", zap.String("url", "www.deepak.com"))
	applicationLogger.Error("this is ERROR message", zap.String("url", "www.deepak.com"))
}
