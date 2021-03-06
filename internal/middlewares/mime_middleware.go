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
	"mime"
	"net/http"

	ierrors "github.com/P4radoX/gomu/internal/errors"
	"github.com/pkg/errors"
)

// MIMEMiddleware middleware function enforces MIME type from Content-Type header and prevents
// sending requests with non-desired types
func MIMEMiddleware(mimeType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" {
			mtype, _, err := mime.ParseMediaType(mimeType)

			// Malformed MIME type
			if err != nil {
				http.Error(w, errors.Wrap(ierrors.ErrMIME, "Corrupted Content-Type header").Error(), http.StatusBadRequest)

				return
			}

			// Not matching MIME type
			if mtype != mimeType {
				http.Error(w, errors.Wrap(ierrors.ErrMIME, "Unsupported Content-Type header value").Error(), http.StatusUnsupportedMediaType)

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
