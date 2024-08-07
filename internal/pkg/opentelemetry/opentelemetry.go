package opentelemetry

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// NewJaegerTraceProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
// Ref: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go#L42
func NewJaegerTraceProvider() (*tracesdk.TracerProvider, error) {
	url := os.Getenv("JAEGER_URL")

	exp, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(url),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create exporter: %w", err)
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(os.Getenv("JAEGER_SERVICE_NAME")),
			attribute.String("environment", os.Getenv("ENVIRONMENT")),
			// attribute.Int64("ID", id),
		)),
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(0.1))), // 10% of traces sampled
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// To trace external requests, we need to use the context propagation
	// https://opentelemetry.io/docs/instrumentation/go/manual/#propagators-and-context
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp, nil
}
