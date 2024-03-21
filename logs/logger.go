package logs

import (
	"os"

	"github.com/rs/zerolog"
)

func CreateLogger() zerolog.Logger {
	println("hello")
	println(os.Getenv("RUN_ENV"))
	if os.Getenv("RUN_ENV") == "test" {
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	}

	logFilePath := "/var/log/myapp/app.log"
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// If opening the log file fails, panic. Alternatively, you could handle this error differently.
		os.Stdout.Write([]byte("Opening file failed"))
		panic(err)
	}

	// For non-test environments, return a logger that writes to both the file and stdout.
	multi := zerolog.MultiLevelWriter(logFile, os.Stdout)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return logger
}
