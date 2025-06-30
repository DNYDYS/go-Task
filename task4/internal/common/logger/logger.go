package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	once   sync.Once
)

// Init 初始化日志器（基础版）
func Init(level zapcore.Level, isJSON bool) {
	once.Do(func() {
		encoder := getEncoder(isJSON)
		writer := zapcore.AddSync(os.Stdout)
		core := zapcore.NewCore(encoder, writer, zap.NewAtomicLevelAt(level))

		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		sugar = logger.Sugar()
	})
}

// InitWithConfig 初始化日志器（高级配置版）
func InitWithConfig(opts ...Option) {
	once.Do(func() {
		cfg := defaultConfig()
		for _, opt := range opts {
			opt(&cfg)
		}

		encoder := getEncoder(cfg.isJSON)
		writer := getWriter(cfg.outputPaths, cfg.errorOutputPaths)
		core := zapcore.NewCore(encoder, writer, cfg.level)

		logger = zap.New(core,
			zap.AddCaller(),
			zap.AddCallerSkip(cfg.callerSkip),
			zap.AddStacktrace(cfg.stacktraceLevel),
		)

		if cfg.addFields != nil {
			logger = logger.With(cfg.addFields...)
		}

		sugar = logger.Sugar()
	})
}

// 获取编码器
func getEncoder(isJSON bool) zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if isJSON {
		return zapcore.NewJSONEncoder(encoderCfg)
	}
	return zapcore.NewConsoleEncoder(encoderCfg)
}

// 获取日志写入器
func getWriter(outputPaths, errorOutputPaths []string) zapcore.WriteSyncer {
	writers := make([]zapcore.WriteSyncer, 0, len(outputPaths))
	for _, path := range outputPaths {
		if path == "stdout" {
			writers = append(writers, zapcore.AddSync(os.Stdout))
		} else if path == "stderr" {
			writers = append(writers, zapcore.AddSync(os.Stderr))
		} else {
			// 这里可以添加文件写入逻辑
			// 实际项目中建议使用 lumberjack 等库实现日志轮转
		}
	}

	if len(writers) == 1 {
		return writers[0]
	}
	return zapcore.NewMultiWriteSyncer(writers...)
}

// L 返回原生 zap.Logger
func L() *zap.Logger {
	if logger == nil {
		panic("logger not initialized")
	}
	return logger
}

// S 返回 SugaredLogger
func S() *zap.SugaredLogger {
	if sugar == nil {
		panic("logger not initialized")
	}
	return sugar
}

// 同步日志缓冲区
func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}
