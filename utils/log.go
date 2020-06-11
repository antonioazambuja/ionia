package utils

import (
	"log"
	"os"
)

// LogOperation - log template
var LogOperation = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
