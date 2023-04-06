package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "PLANTSERVER ", log.Ldate|log.Ltime|log.Lshortfile)
