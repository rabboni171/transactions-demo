package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

func NewLogger(logDir, logFileName string) *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.InfoLevel)

	// Настраиваем формат вывода
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339, // Формат времени
		FullTimestamp:   true,         // Показывать полное время
	})

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// Создаем папку, если она не существует
		err = os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			logger.Warn("Failed to create log directory, using default stderr")
		}
	}

	logFilePath := filepath.Join(logDir, logFileName)

	// Открываем файл для записи логов
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		// Если не удалось открыть файл, выводим лог в консоль
		logger.Warn("Failed to log to file, using default stderr")
	}

	return logger
}
