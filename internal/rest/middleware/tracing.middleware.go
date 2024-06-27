package middleware

import (
	"bytes"

	"io"
	"time"

	"cloud.google.com/go/logging"
	"github.com/gin-gonic/gin"
)

type LoggerProxy interface {
	Get() *logging.Logger
}

type logTraceAuth struct {
	RawBody string            `json:"rawBody"`
	Headers map[string]string `json:"headers"`
}

func logTrace(c *gin.Context, lp LoggerProxy, start time.Time, body string) {
	end := time.Now()
	latency := end.Sub(start)
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	lp.Get().Log(
		logging.Entry{
			Severity: logging.Info,
			HTTPRequest: &logging.HTTPRequest{
				Request:      c.Request,
				RequestSize:  c.Request.ContentLength,
				ResponseSize: int64(c.Writer.Size()),
				Status:       c.Writer.Status(),
				RemoteIP:     c.ClientIP(),
				Latency:      latency,
			},
			Payload: logTraceAuth{
				RawBody: body,
				Headers: headers,
			},
		},
	)
}

func TracingMiddleware(lp LoggerProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		bodyInBytes, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)
		c.Next()
		logTrace(c, lp, start, string(bodyInBytes))
	}
}
