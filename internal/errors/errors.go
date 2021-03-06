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

package errors

import "github.com/pkg/errors"

// ErrUnknown describes an unexpected error
var ErrUnknown = errors.New("unknown error occured")

// ErrSystemNotFound describes a filesystem error when OS can't found requested file or directory
var ErrSystemNotFound = errors.New("no such file or directory")

// ErrSystemIsDirectory describes an unauthorized I/O operation on a directory
var ErrSystemIsDirectory = errors.New("is a directory")

// ErrSystemIO describes an I/O operation error
var ErrSystemIO = errors.New("unable to perform I/O operation")

// ErrSystemUnexpectedEOF describes an unexpected EOF while making I/O operations
var ErrSystemUnexpectedEOF = errors.New("unexpected EOF")

// ErrNetTCPSocket describes a TCP socket error
var ErrNetTCPSocket = errors.New("TCP socket error")

// ErrNetUDPSocket describes an UDP socket error
var ErrNetUDPSocket = errors.New("UDP socket error")

// ErrNetUnixSocket describes an Unix socket error
var ErrNetUnixSocket = errors.New("UNIX socket error")

// ErrNetServerDial describes a failed server connection
var ErrNetServerDial = errors.New("server connection failed")

// ErrDatabaseRequest describes a failed or bad database request 
var ErrDatabaseRequest = errors.New("unable to perform database request")

// ErrHTTP describes a HTTP request error
var ErrHTTP = errors.New("unable to process HTTP request")

// ErrMIME describes a malformed or bad MIME type error
var ErrMIME = errors.New("bad or malformed MIME type")

// ErrJSON describes a JSON marshalling/unmarshalling error
var ErrJSON = errors.New("unable to process JSON object")