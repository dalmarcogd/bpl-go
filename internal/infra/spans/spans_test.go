package spans

import (
	"context"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/dalmarcogd/bpl-go/internal/structs"
	"github.com/dalmarcogd/bpl-go/internal/infra/ctxs"
	"github.com/dalmarcogd/bpl-go/internal/infra/validator"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"testing"
)

func TestSpansService_New(t *testing.T) {
	serviceImpl := New()
	sm := services.New().WithValidator(validator.New()).WithSpans(serviceImpl)

	if err := sm.Init(); err != nil {
		t.Error("unexpected error")
	}
	var err error
	serviceImpl.tracer, err = zipkin.NewTracer(
		http.NewReporter("localhost:8080"),
		zipkin.WithNoopTracer(true),
		zipkin.WithNoopSpan(true),
	)
	if err != nil {
		t.Error(err)
	}

	_, span := sm.Spans().New(ctxs.ContextWithCid(context.Background(), "mycid"), structs.WithOrgId("myorgid"))
	span.Finish()
	if err := serviceImpl.ServiceManager().Close(); err != nil {
		t.Error(err)
	}
}
