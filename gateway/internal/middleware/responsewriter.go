package middleware

import "net/http"

// CustomResponseWriter captures status codes
type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader captures response status
func (w *CustomResponseWriter) WriteHeader(statusCode int) {

	w.StatusCode = statusCode

	w.ResponseWriter.WriteHeader(statusCode)
}