package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func OpentracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()
		var span opentracing.Span

		if parentSpanContext, err := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header)); err != nil {
			span = tracer.StartSpan(r.URL.Path)
		} else {
			span = tracer.StartSpan(r.URL.Path, opentracing.ChildOf(parentSpanContext))
		}

		defer span.Finish()

		ctx := opentracing.ContextWithSpan(r.Context(), span)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func OpentracingMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		var span opentracing.Span

		if parentSpanContext, err := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header)); err != nil {
			span = tracer.StartSpan(c.Request.URL.Path)
		} else {
			span = tracer.StartSpan(c.Request.URL.Path, opentracing.ChildOf(parentSpanContext))
		}

		defer span.Finish()

		ctx := opentracing.ContextWithSpan(c.Request.Context(), span)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
