// MIT License

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

package main

import (
	"github.com/P4radoX/gomu/internal/flags"
)

func main() {
	// Parse micro-service execution flags
	fs := flags.NewFlagSet()

	fs.Add(
		&flags.StringFlag{Name: "bind", Description: "Specify the service addr:port bind", MustBeSet: true, Value: "0.0.0.0:8080"},
		&flags.StringFlag{Name: "endpoint", Description: "Specify the service endpoint URL like /service/v1", MustBeSet: true, Value: ""},
	)

	fs.Parse()

	// Initialize logging facility

	// Initialize router

	// Create & setup view

	// Create & setup controllers

	// Register controllers to router

	// Handle routes & allowed methods

	// Use middlewares

	// Serve HTTP & HTTPS
}
