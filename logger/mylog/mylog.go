package mylog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var time_fomart = "2006-01-02 15:04:05.000"
var time_fomart1 = "20060102150405.000"

// 定义日志结构
type Logger struct {
	filePath string   // 日志输出路径
	level    Loglevel //日志级别
	maxSize  int64    // 文件可保存最大值
}

// 定义级别别名
type Loglevel uint16

// 定义日志级别常量
const (
	INFO Loglevel = iota
	WARN
	DEBUG
)

// 转换日志级别
func parseLogLevel(levelStr string) Loglevel {
	levelStr = strings.ToUpper(levelStr)
	switch levelStr {
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "DEBUG":
		return DEBUG
	default:
		return INFO

	}

}

// 判断是否输出
func (l *Logger) print(lev Loglevel) bool {
	return lev >= l.level
}

// 构造函数
func GetLogger(logLevelStr string) *Logger {
	return &Logger{level: parseLogLevel(logLevelStr)}
}

// 设置日志输出地址
func (l *Logger) SetFilePath(
	path string) {
	l.filePath = path
}
func (l *Logger) SetMaxSiza(size int64) {
	l.maxSize = size
}

//获取方法名 文件地址 行号
func getInfo(n int) (method string, fileName string, lineNo int) {
	pc, fileName, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Print("[日志异常]:%s", "获取行号失败")
		return
	}
	method = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(fileName)
	return method, fileName, lineNo
}

// 警告日志级别
func (l *Logger) Warn(format string, msgInterface ...interface{}) {
	if l.print(WARN) {
		log(l, "WARN", format, msgInterface)
	}

}

//一般信息级别
func (l *Logger) Info(format string, msgInterface ...interface{}) {
	if l.print(INFO) {
		log(l, "info", format, msgInterface)
	}
}

// 调试级别信息
func (l *Logger) Debug(format string, msgInterface ...interface{}) {
	if l.print(DEBUG) {
		log(l, "WARN", format, msgInterface)
	}
}

// 写入日志 控制台输出，如果设置了日志输出地址，则输出到文件
func log(l *Logger, levStr string, format string, msgInterface ...interface{}) {
	msg := fmt.Sprintf(format, msgInterface)
	fmt.Println(msg)
	timeStr := time.Now().Format(time_fomart)
	timeStr1 := time.Now().Format(time_fomart1)
	method, fileName, lineNo := getInfo(2)
	content := fmt.Sprintf("[%s]:%s:%s:%s:%d:%s\n", levStr, timeStr, fileName, method, lineNo, msg)
	fmt.Printf(content)
	if l.filePath != "" {
		if !fileIsExit(l.filePath) {
			os.Mkdir(l.filePath, os.ModePerm)
		}
		disFilePath := path.Join(l.filePath, levStr) //l.filePath + "/" + levStr + "_" + time.Now().Unix()
		//fmt.Println(disFilePath)
		fileInfo, err := os.Stat(disFilePath)
		if err == nil {
			fmt.Println(fileInfo.Size())
			totalSize := fileInfo.Size() + int64(len(content))
			if totalSize > l.maxSize {
				err := os.Rename(disFilePath, fmt.Sprintf("%s_bak_%v", disFilePath, timeStr1))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		fileObj, err := os.OpenFile(disFilePath, os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Sprintf("创建日志失败:%v/n", err)
		}
		defer fileObj.Close()
		fileObj.WriteString(content)
	}

}

// 判断文件是否存在
func fileIsExit(filePath string) bool {
	if filePath != "" {
		_, err := os.Stat(filePath)
		return err == nil || os.IsExist(err)
	}
	return false
}
