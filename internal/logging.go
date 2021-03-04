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

package internal

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// NewLogger function returns a new logging instance with the given
// format (JSON or Text) and default level.
func NewLogger(format string, level log.Level) *log.Logger {
	fieldmap := log.FieldMap{
		log.FieldKeyTime:  "@timestamp",
		log.FieldKeyLevel: "@level",
		log.FieldKeyMsg:   "@message",
		log.FieldKeyFunc:  "@caller",
	}

	logger := &log.Logger{
		Out: os.Stderr,
		ReportCaller: true,
		Level:        level,
	}

	switch format {
	case "json", "JSON", "Json":
		logger.Formatter = &log.JSONFormatter{
			TimestampFormat: time.RFC3339,
			FieldMap: fieldmap,
		}
	default:
		logger.Formatter = &log.TextFormatter{
			TimestampFormat: time.RFC3339,
			DisableLevelTruncation: true,
			FieldMap: fieldmap,
		}
	}

	return logger
}
