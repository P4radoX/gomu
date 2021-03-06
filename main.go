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
	"net/http"

	"github.com/P4radoX/gomu/internal"
	ctl "github.com/P4radoX/gomu/internal/controllers"
	"github.com/P4radoX/gomu/internal/flags"
	mdw "github.com/P4radoX/gomu/internal/middlewares"
	"github.com/P4radoX/gomu/internal/views"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Add application flags
	fs := flags.NewFlagSet()

	fs.Add(
		&flags.StringFlag{Name: "bind", Description: "Specify the service addr:port bind", MustBeSet: true, Value: "0.0.0.0:8080"},
		&flags.StringFlag{Name: "endpoint", Description: "Specify the service endpoint URL like /service/v1", MustBeSet: true, Value: "/app/v1"},
		&flags.BoolFlag{Name: "tls", Description: "Specify if the service must be use HTTP or HTTPS", Value: false},
		&flags.StringFlag{Name: "key", Description: "Specify the service matching private key", Value: ""},
		&flags.StringFlag{Name: "cert", Description: "Specify the server signed certificate", Value: ""},
	)

	// Parse flags with service description
	fs.Parse("Golang micro-service template repository")

	// Initialize logging facility
	logger := internal.NewLogger("json", log.InfoLevel)

	// Initialize router
	R := mux.NewRouter().PathPrefix(fs.Get("endpoint").(*flags.StringFlag).Value).Subrouter()

	// Create & setup view
	healthView := views.NewHealthView("/health", http.MethodGet)
	versionView := views.NewVersionView("/version", http.MethodGet)

	// Create & setup controllers
	healthController := ctl.NewHealthController()
	versionController := ctl.NewVersionController(
		"Gomu",
		"P4radoX",
		"Github",
		"v0.0.0-01ab7f5f9dcc98fbb584c73248ea48ef9c3fd2ea",
		"01ab7f5f9dcc98fbb584c73248ea48ef9c3fd2ea",
		"github.com/P4radoX",
	)

	// Register controllers to router with controllers scope middlewares
	R.Handle(healthView.Path(), mdw.HTTPMethodMiddleware(healthController, healthView.Methods()...))
	R.Handle(versionView.Path(), mdw.HTTPMethodMiddleware(versionController, versionView.Methods()...))

	// Register application-wide middlewares
	R.Use(mdw.LoggingMiddleware(logger))

	// Get application parsed arguments values
	bind := fs.Get("bind").(*flags.StringFlag).Value
	endpoint := fs.Get("endpoint").(*flags.StringFlag).Value
	useTLS := fs.Get("tls").(*flags.BoolFlag).Value
	keyTLS := fs.Get("key").(*flags.StringFlag).Value
	certTLS := fs.Get("cert").(*flags.StringFlag).Value

	// Serve HTTPS or HTTP
	logger.Infof("Now serving at %s%s", bind, endpoint)

	if useTLS {
		logger.Fatal(http.ListenAndServeTLS(bind, certTLS, keyTLS, R))
	} else {
		logger.Fatal(http.ListenAndServe(bind, R))
	}
}
