package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	// sugaredlogger
	sugar := logger.Sugar()
	// 记录一条日志
	sugar.Errorw("这是一个Error",
		"Source", "main.go",
		"错误", "未知错误",
		"Logger", "sugar",
	)
	// normallogger
	logger.Error("这是另一个Error",
		zap.String("Logger", "normal"),
	)
	
	// 存放日志到文件
	flogConf := zap.NewProductionConfig()
	flogConf.OutputPaths = append(flogConf.OutputPaths, "./run.log")
	// 只存放内部错误日志
	flogConf.ErrorOutputPaths = append(flogConf.ErrorOutputPaths, "./run.error.log")
	flogger, _ := flogConf.Build()
	flogger.Info("一条运行记录")
	flogger.Error("一条错误日志")
}
