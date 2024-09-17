package logger

import (
    "time"

    "github.com/sirupsen/logrus"
    "github.com/natefinch/lumberjack"
)

// Logger is a custom logger that writes to a file and supports different log levels
type Logger struct {
    *logrus.Logger
}

// New creates a new Logger instance with file rotation and log level configuration
func New(logFilePath string, maxSizeMB int, maxBackups int, maxAgeDays int, logLevel logrus.Level) *Logger {
    // Create a new Logger instance
    logger := logrus.New()

    // Set log format to JSON for compatibility with ELK
    logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: time.RFC3339,
    })

    // Set log level
    logger.SetLevel(logLevel)

    // Configure log file rotation
    logger.SetOutput(&lumberjack.Logger{
        Filename:   logFilePath,
        MaxSize:    maxSizeMB,  // megabytes
        MaxBackups: maxBackups,
        MaxAge:     maxAgeDays, // days
    })

    return &Logger{logger}
}

