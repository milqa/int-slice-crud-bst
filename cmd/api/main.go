package main

import (
	"log"
	"net/http"

	"github.com/milQA/int-slice-crud-bst/internal/api/server"
	"github.com/milQA/int-slice-crud-bst/internal/api/store/bstStore"
	"github.com/milQA/int-slice-crud-bst/pkg/logger"
	"go.uber.org/zap/zapcore"
)

func main() {

	logger, err := logger.NewLogger(zapcore.DebugLevel)
	if err != nil {
		log.Fatalf("cannot init logger: %s", err.Error())
	}
	store := bstStore.NewStore(logger)

	srv := server.NewServer(logger, store)

	logger.Info("start server http://localhost:8080")
	logger.Fatal(http.ListenAndServe(":8080", srv).Error())
}
