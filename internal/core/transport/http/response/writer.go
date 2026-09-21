package core_http_response

import "net/http"

const StatusCodeUninitialized = -1

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitialized,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode != StatusCodeUninitialized {
		return
	}

	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	if w.statusCode == StatusCodeUninitialized {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriter) GetStatusCode() int {
	if w.statusCode == StatusCodeUninitialized {
		panic("no status code set")
	}

	return w.statusCode
}
