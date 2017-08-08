package main

func main() {

}

package http

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (fn HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	fn(w, r)
}



type Middleware interface {
	WrapHandler(Handler) Handler
}

type DDMetricsBackend struct {
	creds string
	metrics []*MetricDef
	prefix string
	...
}

func (m DDMetricsBackend) WrapHandler(h Handler) Handler {
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		for _, m := range m.metrics {
			m.applyMetricToRequest(r)
		}
		h.ServeHTTP(w, r)
	}
}

type HandlerWithMiddleware struct {
	h Handler
	ms []Middleware
}

func (h HandlerWithMiddleware) ServeHTTP(w ResponseWriter, r *Request) {
	for _, mw := range ms {
		h = mw.WrapHandler(h)
	}
	return h.ServeHTTP(w, r)
}

