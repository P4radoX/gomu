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

package controllers

import (
	"fmt"
	"net/http"
)

// VersionController struct represents the /version sub-endpoint controller
type VersionController struct {
	Name       string
	Maintainer string
	Provider   string
	Tag        string
	Commit     string
	URL        string
}

// NewVersionController function returns a new VersionController struct pointer
func NewVersionController(name, maintainer, provider, tag, commit, url string) *VersionController {
	return &VersionController{
		Name:       name,
		Maintainer: maintainer,
		Provider:   provider,
		Tag:        tag,
		Commit:     commit,
		URL:        url,
	}
}

func (ctl *VersionController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Write HTTP headers
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write HTTP response status
	w.WriteHeader(http.StatusOK)

	// Write payload
	w.Write([]byte(
		fmt.Sprintf("{\"name\":\"%s\", \"maintainer\":\"%s\", \"provider\":\"%s\", \"tag\":\"%s\", \"commit\":\"%s\", \"url\":\"%s\"}", ctl.Name, ctl.Maintainer, ctl.Provider, ctl.Tag, ctl.Commit, ctl.URL),
	))
}
