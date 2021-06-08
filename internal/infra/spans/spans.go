package spans

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/infra/ctxs"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/dalmarcogd/bpl-go/internal/structs"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"runtime"
	"strings"
	"time"
)

type (
	spansService struct {
		services.NoopHealth
		sis         services.Sis
		ctx         context.Context
		tracer      *zipkin.Tracer
		reporter    reporter.Reporter
		host        string
		serviceName string
		version     string
	}
)

func New() *spansService {
	return &spansService{}
}

func (s *spansService) Init(ctx context.Context) error {
	s.ctx = ctx
	if s.tracer == nil {
		if s.host == "" {
			s.host = s.Sis().Environment().SpanUrl()
		}
		s.reporter = http.NewReporter(s.host, http.BatchInterval(time.Second*3))

		if s.serviceName == "" {
			s.serviceName = s.Sis().Environment().Service()
		}
		// create our local spansService endpoint
		endpoint, err := zipkin.NewEndpoint(s.serviceName, "0.0.0.0:8080")
		if err != nil {
			return err
		}

		s.tracer, err = zipkin.NewTracer(
			s.reporter,
			zipkin.WithLocalEndpoint(endpoint),
			zipkin.WithTraceID128Bit(true),
		)
		if err != nil {
			return err
		}
	}

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(zipkintracer.Wrap(s.tracer))
	return nil
}

func (s *spansService) Close() error {
	if err := s.reporter.Close(); err != nil {
		return err
	}
	return nil
}

func (s *spansService) WithSis(c services.Sis) services.Spans {
	s.sis = c
	return s
}

func (s *spansService) Sis() services.Sis {
	return s.sis
}

func (s *spansService) New(ctx context.Context, spanConfigs ...structs.SpanConfig) (context.Context, *structs.Span) {
	var cid string
	if c := ctxs.GetCidFromContext(ctx); c != nil {
		cid = *c
	}

	funcName := ""
	line := 0
	fileName := ""
	if pc, f, l, ok := runtime.Caller(1); ok {
		funcName = runtime.FuncForPC(pc).Name()
		lastDot := strings.LastIndexByte(funcName, '.')
		if lastDot < 0 {
			lastDot = 0
		}
		funcName = funcName[lastDot+1:]
		fileName = f
		line = l
	}

	sp := &structs.Span{
		Name:         funcName,
		Cid:          cid,
		Line:         line,
		FileName:     fileName,
		FuncName:     funcName,
		Version:      s.version,
		Resource:     "cpu",
		Custom:       map[string]interface{}{},
		InternalSpan: nil,
	}

	for _, config := range spanConfigs {
		config.Apply(ctx, sp)
	}

	sp.InternalSpan, ctx = s.tracer.StartSpanFromContext(ctx, sp.Name)

	return ctx, sp
}

func (s *spansService) Tracer() *zipkin.Tracer {
	return s.tracer
}
