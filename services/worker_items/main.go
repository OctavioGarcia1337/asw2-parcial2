package main

import (
	log "github.com/sirupsen/logrus"
	worker "worker/controllers/worker"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Starting worker")
	worker.StartWorker()
}
