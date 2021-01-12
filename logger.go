/**
* @Author: Aaron
* @Date: 2020/11/6 16:48
 */
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	Info           *log.Logger
	Warn           *log.Logger
	Error          *log.Logger
	FileNamePrefix string
	errFile        *os.File
	infoFile       *os.File
	warnFile       *os.File
}

func NewLog(logfileName string) *Logger {
	logger := Logger{
		FileNamePrefix: logfileName,
	}
	if !logger.Init() {
		return nil
	}
	return &logger
}

func (logger *Logger) Init() bool {
	path := "./log"
	ex, err := PathExists(path)
	if !ex {
		//创建目录
		//perm权限设置，os.ModePerm为0777
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}

	errFile, err := os.OpenFile(fmt.Sprintf("%s/errors_%s.log", path, logger.FileNamePrefix), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
		return false
	}
	infoFile, err := os.OpenFile(fmt.Sprintf("./log/info_%s.log", logger.FileNamePrefix), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
		return false
	}
	warnFile, err := os.OpenFile(fmt.Sprintf("./log/warn_%s.log", logger.FileNamePrefix), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
		return false
	}

	logger.Info = log.New(io.MultiWriter(os.Stdout, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warn = log.New(io.MultiWriter(os.Stdout, warnFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	// io.MultiWriter函数可以包装多个io.Writer为一个io.Writer，这样我们就可以达到同时对多个io.Writer输出日志的目的。
	logger.Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

	logger.infoFile = infoFile
	logger.errFile = errFile
	logger.warnFile = warnFile
	return true
}

//func (logger *Logger) info(v ...interface{}) {
//	logger.Info.Println(v...)
//}
//
//func (logger *Logger) warn(v ...interface{}) {
//	logger.Warn.Println(v...)
//}
//
//func (logger *Logger) error(v ...interface{}) {
//	logger.Error.Println(v...)
//}

func (logger *Logger) close() {
	logger.infoFile.Close()
	logger.errFile.Close()
	logger.warnFile.Close()
}
