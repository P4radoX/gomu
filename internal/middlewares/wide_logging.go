// Copyright (c) 2021 P4radoX

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package middlewares

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// LoggingMiddleware application-wide middleware logs every HTTP request
func LoggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			wrw := wrapResponseWriter(w)

			start := time.Now()
			next.ServeHTTP(wrw, r)
			t := time.Now()

			logger.WithFields(log.Fields{
				"@protocol_version": r.Proto,
				"@method":           r.Method,
				"@request_uri":      r.URL.EscapedPath(),
				"@remote_addr":      r.RemoteAddr,
				"@status":           wrw.Status(),
				"@duration":         t.Sub(start).Milliseconds(),
			}).Info("New incoming service request")
		}

		return http.HandlerFunc(fn)
	}
}
