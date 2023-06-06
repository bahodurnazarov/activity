package utils

import (
	log "github.com/sirupsen/logrus"
	logg "log"
	"os"
)

var (
	Server *logg.Logger
	Errl   *log.Logger

	info *os.File
	errl *os.File
)

func createFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func init() {
	info = createFile("../../pkg/logs/server.log")
	errl = createFile("../../pkg/logs/error.log")

	Errl = log.New()
	Server = logg.New(info, "INFO: ", logg.Ldate|logg.Ltime|logg.Lshortfile)

	Errl.SetOutput(errl)
}

func CloseLogger() {
	err := info.Close()
	if err != nil {
		log.Fatalf("can't close file %v, %v", info, err)
	}
	err = errl.Close()
	if err != nil {
		log.Fatalf("can't close file %v, %v", errl, err)
	}
}
