package logger

import "go.uber.org/zap/zapcore"

type Config struct {
	level            zapcore.Level
	isJSON           bool
	outputPaths      []string
	errorOutputPaths []string
	callerSkip       int
	stacktraceLevel  zapcore.Level
	addFields        []zapcore.Field
}

type Option func(*Config)

func defaultConfig() Config {
	return Config{
		level:            zapcore.InfoLevel,
		isJSON:           true,
		outputPaths:      []string{"stdout"},
		errorOutputPaths: []string{"stderr"},
		callerSkip:       1,
		stacktraceLevel:  zapcore.ErrorLevel,
	}
}

// WithLevel 设置日志级别
func WithLevel(level zapcore.Level) Option {
	return func(c *Config) {
		c.level = level
	}
}

// WithJSONFormat 设置JSON格式
func WithJSONFormat(enabled bool) Option {
	return func(c *Config) {
		c.isJSON = enabled
	}
}

// WithOutputPaths 设置输出路径
func WithOutputPaths(paths ...string) Option {
	return func(c *Config) {
		c.outputPaths = paths
	}
}

// WithErrorOutputPaths 设置错误输出路径
func WithErrorOutputPaths(paths ...string) Option {
	return func(c *Config) {
		c.errorOutputPaths = paths
	}
}

// WithCallerSkip 设置调用者跳过层数
func WithCallerSkip(skip int) Option {
	return func(c *Config) {
		c.callerSkip = skip
	}
}

// WithStacktraceLevel 设置堆栈跟踪级别
func WithStacktraceLevel(level zapcore.Level) Option {
	return func(c *Config) {
		c.stacktraceLevel = level
	}
}

// WithFields 添加默认字段
func WithFields(fields ...zapcore.Field) Option {
	return func(c *Config) {
		c.addFields = fields
	}
}
