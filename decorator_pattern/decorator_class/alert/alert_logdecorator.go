package alert

import (
	log "github.com/sirupsen/logrus"
)

// LogrusLogger :
type LogrusLogger struct {
}

// Log :
func (ll LogrusLogger) Log(msg string) {
	log.WithFields(log.Fields{"msg": msg}).Info("A Logrus Log")
}

// AlertLogDecorator :
type AlertLogDecorator struct {
	LogClient   Logger
	AlertClient Alert
}

// Notify :
func (ald AlertLogDecorator) Notify(msg string) {
	ald.LogClient.Log(msg)
	ald.AlertClient.Notify(msg)
}
