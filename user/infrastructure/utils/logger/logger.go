package logger

import (
    log "github.com/sirupsen/logrus"
)

func Info(message string) {
	log.Info(message)
}
