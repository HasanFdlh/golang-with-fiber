package config

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitFiberLogger() logger.Config {
	// bikin nama file sesuai tanggal
	logFile := "app-" + time.Now().Format("2006-01-02") + ".log"

	// open file
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// gabungkan terminal + file
	mw := io.MultiWriter(os.Stdout, file)

	return logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     mw,
	}
}
