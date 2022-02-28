// init.go
//
// Aternity Tech-Community
// 109-opentelemetry-export
// version: 22.3.1
//
// Patch for example app main.go to use environmnent variable

package main

import (
	"os"
)
var (
	OTEL_EXPORTER_OTLP_ENDPOINT = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
)