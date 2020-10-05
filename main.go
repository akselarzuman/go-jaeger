package main

import (
	"log"
	"net/http"

	"github.com/akselarzuman/go-jaeger/configuration"
	"github.com/akselarzuman/go-jaeger/jaegerwrapper"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	configuration.InitializeEnv()
	tracer, closer := jaegerwrapper.NewFromEnv()
	defer (*closer).Close()

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		// tracer := opentracing.GlobalTracer()

		// Extract the context from the headers
		spanCtx, _ := (*tracer).Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		serverSpan := (*tracer).StartSpan("server", ext.RPCServerOption(spanCtx))
		defer serverSpan.Finish()
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
