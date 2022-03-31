package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/akselarzuman/go-jaeger/internal/pkg/opentelemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// this snippet makes an HTTP request to the API
func main() {
	tp, err := opentelemetry.NewJaegerTraceProvider()
	if err != nil {
		log.Fatal(err.Error())
	}

	if tp != nil {
		defer func() {
			// Cleanly shutdown and flush telemetry when the application exits.
			exitCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := tp.Shutdown(exitCtx); err != nil {
				log.Println(err.Error())
			}
		}()
	}

	ctx, span := opentelemetry.NewSpan(context.Background(), "httpRequest")
	defer span.End()

	opentelemetry.AddSpanEvents(span, "client event", map[string]string{
		"event": "client event",
	})

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/health_check", nil)
	if err != nil {
		opentelemetry.AddSpanError(span, err)
		// trace.FailSpan(span, "request error")
		log.Fatal(err.Error())
	}

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	res, err := client.Do(req)
	if err != nil {
		opentelemetry.AddSpanError(span, err)
		// trace.FailSpan(span, "rounttrip error")

		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("unexpected response code")
		opentelemetry.AddSpanError(span, err)
		// trace.FailSpan(span, "response error")

		log.Fatal(err.Error())
	}

	log.Println("response:", res.StatusCode)
}
