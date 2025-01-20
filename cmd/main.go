package main

import (
	"github.com/jmechavez/EmailAudit/infrastructure/http"
	"github.com/jmechavez/EmailAudit/infrastructure/logger"
)

func main() {
	logger.Info("Starting the application")
	http.Start()
}
