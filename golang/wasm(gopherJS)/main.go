package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall/js"
	"time"
)

var (
	err     error
	logFile *os.File
)

func main() {
	if logFile, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666); err != nil {
		log.Println("Failed to open or create log file: " + err.Error())
		log.SetOutput(os.Stdout)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Failed to start server: " + err.Error())
	}
	select {}
}