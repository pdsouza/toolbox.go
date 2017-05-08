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
	"fmt"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	var cases = []struct {
		URL string
	}{
		{"https://github.com/maruos/maruos/releases/download/v0.4/maru-v0.4-sha1sums.asc.txt"},
		{"https://github.com/maruos/blueprints/releases/download/v0.4/maru-v0.4-jessie-rootfs-b1656e03.tar.gz"},
	}

	for i, tt := range cases {
		req, err := NewDownloadRequest(tt.URL)
		if err != nil {
			t.Errorf("case #%d: Failed to create download request: %s", i, err.Error())
		}

		gotStart, gotEnd := false, false
		req.ProgressHandler = func(percent float64) {
			if percent == 0 {
				gotStart = true
			} else if percent == 1 {
				gotEnd = true
			}
		}

		path, err := req.Download()
		if err != nil {
			t.Errorf("case #%d: Failed to download: %s", i, err.Error())
		}

		if !gotStart || !gotEnd {
			t.Errorf("case #%d: Failed to get download progress. gotStart: %s, gotEnd: %s", i, gotStart, gotEnd)
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("case #%d: Failed to find downloaded file: %s", i, err.Error())
		}

		// clean up
		os.Remove(path)
	}
}
