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

import "net/http"

// wrappedResponseWriter struct represents a wrapped http.ResponseWriter interface
// It can be used to get the status code after the service has wrote a response.
// Otherwise, use directly the middleware http.ResponseWriter
type wrappedResponseWriter struct {
	http.ResponseWriter
	status         int
	hasWroteHeader bool
}

// wrapResponseWriter wraps an http.ResponseWriter and returns a new wrappedResponseWriter struct pointer
func wrapResponseWriter(w http.ResponseWriter) *wrappedResponseWriter {
	return &wrappedResponseWriter{ResponseWriter: w}
}

// Status method returns the wrapped http.ResponseWriter status code
func (wrw *wrappedResponseWriter) Status() int {
	return wrw.status
}

// WriteHeader method satisfy http.ResponseWriter interface
func (wrw *wrappedResponseWriter) WriteHeader(code int) {
	if wrw.hasWroteHeader {
		return
	}

	wrw.status = code
	wrw.ResponseWriter.WriteHeader(code)
	wrw.hasWroteHeader = true

	return
}
