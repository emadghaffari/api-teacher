package jgr

import (
	"fmt"
	"io"
	"sync"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

var (
	once sync.Once

	// Tracing var
	Tracing Tracer = &tracing{}
)

// Tracer interface for jaeger
type Tracer interface {
	New(service string)
	GetCloser() io.Closer
	GetTracer() opentracing.Tracer
}

type tracing struct {
	closer io.Closer
	tracer opentracing.Tracer
}

func (t *tracing) New(service string) {
	once.Do(func() {
		cfg := jaegercfg.Configuration{
			ServiceName: service,
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: "jaeger:6831",
			},
		}

		// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
		// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
		// frameworks.
		jLogger := jaegerlog.StdLogger
		jMetricsFactory := metrics.NullFactory

		// Initialize tracer with a logger and a metrics factory
		tracer, closer, _ := cfg.NewTracer(
			jaegercfg.Logger(jLogger),
			jaegercfg.Metrics(jMetricsFactory),
		)

		opentracing.SetGlobalTracer(tracer)
		t.closer = closer
		t.tracer = tracer

		fmt.Println("jaeger connected")
	})
}

func (t *tracing) GetCloser() io.Closer {
	return t.closer
}

func (t *tracing) GetTracer() opentracing.Tracer {
	return t.tracer
}
