package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/milQA/int-slice-crud-bst/internal/api/server"
	"github.com/milQA/int-slice-crud-bst/internal/api/service"
	"github.com/milQA/int-slice-crud-bst/internal/api/store/bstStore"
	"github.com/milQA/int-slice-crud-bst/pkg/logger"
)

var (
	port        string
	host        string
	loggerLevel string

	intSliceUploadFilePath string
)

const (
	defaultPort        = "8080"
	defaultHost        = "0.0.0.0"
	defaultLoggerLevel = "DEBUG"

	defaultIntSliceUploadFilePath = "./backup.json"
)

func main() {

	flag.StringVar(&port, "port", defaultPort, "server port")
	flag.StringVar(&host, "host", defaultHost, "server host")
	flag.StringVar(&loggerLevel, "loglvl", defaultLoggerLevel, "set logger level")
	flag.StringVar(&intSliceUploadFilePath, "upload",
		defaultIntSliceUploadFilePath, "path file to upload int slice")

	flag.Parse()

	logger, err := logger.NewLogger(loggerLevel)
	if err != nil {
		log.Fatalf("cannot init logger: %s", err.Error())
	}

	store := bstStore.NewStore(logger)

	service, err := service.NewService(logger, store,
		service.WithUpdateIntSliceByFile(intSliceUploadFilePath))
	if err != nil {
		logger.Sugar().Fatalf("cannot init service: %s", err.Error())
	}

	srv := server.NewServer(logger, service)

	log.Print(fmt.Sprintf("start server http://%s:%s logger level: %s", host, port, loggerLevel))

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf("%s:%s", host, port), srv).Error(),
	)
}
