package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志格式 决定日志的输出格式
func getEncoder() zapcore.Encoder {
	//生成环境默认配置
	config := zap.NewProductionEncoderConfig()
	// 配置时间格式 日志级别格式 代码调用位置
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回配置后的JSON格式编码器实例
	return zapcore.NewJSONEncoder(config)
}

// 日志写入器 决定日志的输出位置
func getWriterSyncer() zapcore.WriteSyncer {
	// 在控制台和文件中写入日志
	lumberJack := &lumberjack.Logger{
		Filename:   "./log/app.log",
		MaxAge:     30,   // 单个文件的保留天数
		MaxSize:    100,  // 单个文件的最大大小MB
		Compress:   true, // 是否压缩日志文件
		MaxBackups: 5,    // 最大备份文件数
	}

	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),  // 控制台
		zapcore.AddSync(lumberJack), // 文件
	)
}

// 异步日志 防止日志写入阻塞主程序执行
func wrapAsyncCore(ws zapcore.WriteSyncer) *zapcore.BufferedWriteSyncer {
	asyncWS := &zapcore.BufferedWriteSyncer{
		WS:            ws,
		Size:          256 * 1024,
		FlushInterval: time.Second * 30,
		Clock:         zapcore.DefaultClock,
	}

	return asyncWS
}

// 日志采样
func wrapSampler(core zapcore.Core) zapcore.Core {
	return zapcore.NewSampler(core, time.Second, 100, 100)
}

// 将上面组件组装成完整的zap日志
func newZapLogger() *zap.Logger {
	// 日志格式 编码器
	encoder := getEncoder()
	// 日志写入器
	ws := getWriterSyncer()
	// 日志级别 打印INFO及以上日志
	level := zapcore.InfoLevel

	// 组装核心 异步日志直接写入核心中
	core := zapcore.NewCore(encoder, wrapAsyncCore(ws), level)

	// 异步 采用 调用者 错误堆栈
	logger := zap.New(
		core,
		zap.WrapCore(wrapSampler),         // 开启采样
		zap.AddCaller(),                   // 开启调用者信息
		zap.AddStacktrace(zap.ErrorLevel), // 开启错误堆栈
	)

	return logger
}

var Logger *zap.Logger

// 初始化日志实例
func InitLogger() {
	Logger = newZapLogger()
}
