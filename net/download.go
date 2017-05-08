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
	ProgressHandler DownloadProgressHandler
}

func NewDownloadRequest(url string) (*DownloadRequest, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return &DownloadRequest{req, nil}, nil
}

type DownloadProgressHandler func(percent float64)

func (d *DownloadRequest) Download() (path string, err error) {
	path = ""
	resp, err := http.DefaultClient.Do(d.Request)
	if err != nil {
		return path, err
	}
	defer resp.Body.Close()

	tokens := strings.Split(d.Request.URL.String(), "/")
	path = tokens[len(tokens)-1]
	partialFile := path + ".partial"
	fout, err := os.Create(partialFile)
	if err != nil {
		return path, err
	}

	if err := copyAndReportProgress(fout, resp.Body, resp.ContentLength, d.ProgressHandler); err != nil {
		fout.Close()
		return path, err
	}
	fout.Close()

	if err := os.Rename(partialFile, path); err != nil {
		return path, err
	}

	return path, nil
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
