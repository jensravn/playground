package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

func init() {
	// Disable log prefixes such as the default timestamp.
	// Prefix text prevents the message from being parsed as JSON.
	// A timestamp is added when shipping logs to Cloud Logging.
	log.SetFlags(0)
}

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

func projectID() string {
	projectID, err := metadata.ProjectID()
	if err != nil {
		log.Printf("metadata.ProjectID(): %v", err)
		projectID = ""
	}
	log.Printf("project id is %s", projectID)
	return projectID
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logInfo(r, "info message")
	fmt.Fprintln(w, "Hello info!")
	logWarn(r, "warning message")
	fmt.Fprintln(w, "Hello warning!")
	logError(r, "error message")
	fmt.Fprintln(w, "Hello error!")
}

func logError(r *http.Request, msg string) {
	logFn(r, msg, "ERROR")
}
func logInfo(r *http.Request, msg string) {
	logFn(r, msg, "INFO")
}
func logWarn(r *http.Request, msg string) {
	logFn(r, msg, "WARNING")
}

func logFn(r *http.Request, msg, severity string) {
	projectID := projectID()
	var trace string
	if projectID != "" {
		traceHeader := r.Header.Get("X-Cloud-Trace-Context")
		log.Println(fmt.Sprintf("traceHeader: %s", traceHeader))
		traceParts := strings.Split(traceHeader, "/")
		log.Println(fmt.Sprintf("traceParts: %v", traceParts))
		if len(traceParts) > 0 && len(traceParts[0]) > 0 {
			trace = fmt.Sprintf("projects/%s/traces/%s", projectID, traceParts[0])
		}
	}
	log.Println(fmt.Sprintf("trace: %s", trace))
	log.Println(Entry{
		Severity: severity,
		Message:  msg,
		// Component: "arbitrary-property",
		Trace: trace,
	})
}

// Entry defines a log entry.
type Entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity,omitempty"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"`

	// Logs Explorer allows filtering and display of this as `jsonPayload.component`.
	Component string `json:"component,omitempty"`
}

// String renders an entry structure to the JSON format expected by Cloud Logging.
func (e Entry) String() string {
	if e.Severity == "" {
		e.Severity = "INFO"
	}
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}
