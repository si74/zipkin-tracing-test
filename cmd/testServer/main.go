package main

import (
	"io"
	"log"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	// addr := os.Getenv("ZIPKIN_ADDR")
	// if addr == "" {
	// 	log.Fatal("zipkin address not set")
	// }

	// serviceName := "testserver"
	// debug := false
	// traceID128Bit := true

	// create collector.
	// collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	// if err != nil {
	// 	log.Fatalf("unable to create Zipkin HTTP collector: %+v\n", err)
	// }
	//
	// // create recorder.
	// recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)
	//
	// // create tracer.
	// tracer, err := zipkin.NewTracer(
	// 	recorder,
	// 	zipkin.ClientServerSameSpan(sameSpan),
	// 	zipkin.TraceID128Bit(traceID128Bit),
	// )
	// if err != nil {
	// 	log.Fatalf("unable to create Zipkin tracer: %+v\n", err)
	// }

	//opentracing.InitGlobalTracer(tracer)

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"your_service_name",
		config.Logger(jaeger.StdLogger),
	)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// var sp opentracing.Span
		// opName := req.URL.Path
		// wireContext, err := opentracing.GlobalTracer().Extract(
		// 	opentracing.HTTPHeaders,
		// 	opentracing.HTTPHeadersCarrier(req.Header))
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//
		// sp = opentracing.StartSpan(
		// 	opName,
		// 	ext.RPCServerOption(wireContext))
		// defer sp.Finish()
		parent := opentracing.GlobalTracer().StartSpan("hello")
		defer parent.Finish()
		child := opentracing.GlobalTracer().StartSpan(
			"world", opentracing.ChildOf(parent.Context()))
		defer child.Finish()

		//ctx := opentracing.ContextWithSpan(context.Background(), sp)

		io.WriteString(w, "salut! ca va?")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
