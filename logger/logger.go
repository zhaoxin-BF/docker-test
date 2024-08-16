package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type LogEntry struct {
	Level  string    `json:"level"`
	Ts     time.Time `json:"ts"`
	Caller string    `json:"caller"`
	Msg    string    `json:"msg"`
}

type Logger struct {
	w          *log.Logger
	logDirPath string
}

func NewLogger(logDirPath string) *Logger {
	if err := os.MkdirAll(logDirPath, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logFilePath := filepath.Join(logDirPath, fmt.Sprintf("everai-%s.log", time.Now().Format("20060102")))
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(logFile, os.Stdout)

	return &Logger{
		w:          log.New(multiWriter, "", 0),
		logDirPath: logDirPath,
	}
}

func getLastTwoComponents(input string) string {
	// 按 "/" 切分字符串
	parts := strings.Split(input, "/")

	// 获取最后两个组件
	lastComponents := parts[len(parts)-2] + "/" + parts[len(parts)-1]

	return lastComponents
}

func (l *Logger) logEntry(level, caller, msg string) {
	caller = getLastTwoComponents(caller)
	entry := LogEntry{
		Level:  level,
		Ts:     time.Now(),
		Caller: caller,
		Msg:    msg,
	}

	entryJSON, err := json.Marshal(entry)
	if err != nil {
		l.w.Printf("Failed to marshal log entry: %v", err)
		return
	}

	l.w.Println(string(entryJSON))
}

func (l *Logger) Debug(msg string) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("debug", fmt.Sprintf("%s:%d", file, line), msg)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("debug", fmt.Sprintf("%s:%d", file, line), fmt.Sprintf(format, v...))
}

func (l *Logger) Info(msg string) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("info", fmt.Sprintf("%s:%d", file, line), msg)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("info", fmt.Sprintf("%s:%d", file, line), fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(msg string) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("warn", fmt.Sprintf("%s:%d", file, line), msg)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("warn", fmt.Sprintf("%s:%d", file, line), fmt.Sprintf(format, v...))
}

func (l *Logger) Error(msg string) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("error", fmt.Sprintf("%s:%d", file, line), msg)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("error", fmt.Sprintf("%s:%d", file, line), fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(msg string) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("fatal", fmt.Sprintf("%s:%d", file, line), msg)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logEntry("fatal", fmt.Sprintf("%s:%d", file, line), fmt.Sprintf(format, v...))
	os.Exit(1)
}
