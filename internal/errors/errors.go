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

// UnknownError describes an unexpected error
var UnknownError error

// SystemNotFoundError describes a filesystem error when OS can't found requested file or directory
var SystemNotFoundError error

// SystemIsDirectoryError describes an unauthorized I/O operation on a directory
var SystemIsDirectoryError error

// SystemIOError describes an I/O operation error
var SystemIOError error

// SystemUnexpectedEOFError describes an unexpected EOF while making I/O operations
var SystemUnexpectedEOFError error

// NetTCPSocketError describes a TCP socket error
var NetTCPSocketError error

// NetUDPSocketError describes an UDP socket error
var NetUDPSocketError error

// NetUnixSocketError describes an Unix socket error
var NetUnixSocketError error

// NetServerDialError describes a failed server connection
var NetServerDialError error

// DatabaseRequestError describes a failed or bad database request 
var DatabaseRequestError error

// HTTPBadRequestQueryError describes a bad URL query request
var HTTPBadRequestQueryError error

// HTTPMIMETypeError describes a malformed or bad MIME type error
var HTTPMIMETypeError error

// JSONMarshalError describes a JSON marshalling/unmarshalling error
var JSONMarshalError error