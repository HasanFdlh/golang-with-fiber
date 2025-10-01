package config

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitFiberLogger() logger.Config {
	// bikin folder logs kalau belum ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatalf("❌ gagal bikin folder logs: %v", err)
		}
	}

	// bikin nama file sesuai tanggal
	logFile := "logs/app-" + time.Now().Format("2006-01-02") + ".log"

	// open file (append + write only)
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("❌ error opening log file: %v", err)
	}

	// gabungkan output ke terminal + file
	mw := io.MultiWriter(os.Stdout, file)

	return logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     mw,
	}
}
