//
// Copyright 2017 Preetam J. D'Souza
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package net

import (
	"io"
	"net/http"
	"os"
	"strings"
)

type DownloadRequest struct {
	Request         *http.Request
	Filename        string
	ProgressHandler DownloadProgressHandler
}

func NewDownloadRequest(url string) (*DownloadRequest, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// assume last element of path is download name
	tokens := strings.Split(url, "/")
	name := tokens[len(tokens)-1]

	return &DownloadRequest{req, name, nil}, nil
}

type DownloadProgressHandler func(percent float64)

func (d *DownloadRequest) Download() (path string, err error) {
	resp, err := http.DefaultClient.Do(d.Request)
	if err != nil {
		return d.Filename, err
	}
	defer resp.Body.Close()

	partialFile := d.Filename + ".partial"
	fout, err := os.Create(partialFile)
	if err != nil {
		return d.Filename, err
	}

	if err := copyAndReportProgress(fout, resp.Body, resp.ContentLength, d.ProgressHandler); err != nil {
		fout.Close()
		return d.Filename, err
	}
	fout.Close()

	if err := os.Rename(partialFile, d.Filename); err != nil {
		return d.Filename, err
	}

	return d.Filename, nil
}

func copyAndReportProgress(dst io.Writer, src io.Reader, length int64, handler DownloadProgressHandler) error {
	if handler == nil || length < 0 {
		_, err := io.Copy(dst, src)
		return err
	}

	var totalWritten int64 = 0

	// send 0% before we start
	handler(float64(totalWritten) / float64(length))

	for {
		written, err := io.CopyN(dst, src, 32*1024)
		totalWritten += written
		handler(float64(totalWritten) / float64(length))
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}

	return nil
}
