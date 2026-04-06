package telemetry

import (
	"context"

	"github.com/m-bromo/rolldeck/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span)
	Shutdown(ctx context.Context) error
}

type Span interface {
	End(options ...trace.SpanEndOption)
	SetAttributes(attrs ...attribute.KeyValue)
}

type OpenTelemetryTracer struct {
	tracer   trace.Tracer
	provider *sdktrace.TracerProvider
}

func NewOpenTelemetryTracer(cfg *config.Config) (Tracer, error) {
	ctx := context.Background()

	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(cfg.Telemetry.Endpoint),
	}

	exporter, err := otlptracehttp.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(30)),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Telemetry.ServiceName),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return &OpenTelemetryTracer{
		tracer:   tracerProvider.Tracer(cfg.Telemetry.ServiceName),
		provider: tracerProvider,
	}, nil
}

func (t *OpenTelemetryTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, spanName, opts...)
}

func (t *OpenTelemetryTracer) Shutdown(ctx context.Context) error {
	return t.provider.Shutdown(ctx)
}
