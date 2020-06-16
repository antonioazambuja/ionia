package utils

import (
	"net/http"
	"log"
	"os"
)

// LogOperation - log template
var LogOperation = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

// ServiceLog - Log for any requested
func ServiceLog(statusCode int, r *http.Request, serviceName string) {
	LogOperation.Printf("Status: %d - Service: %s - Endpoint: %s - %s, %s, %s", statusCode, serviceName, r.RequestURI, r.UserAgent(), r.Host, r.Method)
}
