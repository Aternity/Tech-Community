# Dockerfile 
#
# Aternity Tech-Community
# 109-opentelemetry-export
# version: 22.3.1
#
# Example
#
# - Build
#
#   docker build -t local/service109-go:latest .
#
# - Run
#
#   docker run  -it --rm -p 8109:80 -e OTEL_EXPORTER_OTLP_ENDPOINT="otel-collector:4317" -e OTEL_SERVICE_NAME="service109-go" local/service109-go:latest
#
# https://hub.docker.com/_/golang
FROM golang:1.17

WORKDIR /usr/src/example

# Get OpenTelemetry example main.go
RUN curl -Ol https://raw.githubusercontent.com/open-telemetry/opentelemetry-go/main/example/otel-collector/main.go

########################################################################
# Patch example app main.go to use environmnent variable
# Local context
COPY init.go .
# TODO: remote context
# Apply patch
RUN sed -i 's/\"localhost:30080\"/OTEL_EXPORTER_OTLP_ENDPOINT/g' main.go
########################################################################

RUN go mod init example/main
RUN go mod tidy 

# RUN go install ./main.go

CMD [ "go" , "run" , "init.go" , "main.go" ]