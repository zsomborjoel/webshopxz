package middleware

import "github.com/gin-gonic/gin"

type responseWriterWithInterceptor struct {
	gin.ResponseWriter
	responseBody []byte
}

func (w *responseWriterWithInterceptor) Write(data []byte) (int, error) {
	w.responseBody = append(w.responseBody, data...)
	return w.ResponseWriter.Write(data)
}

func (w *responseWriterWithInterceptor) Body() string {
	return string(w.responseBody)
}
