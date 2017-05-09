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

package ui

import "testing"

func TestRender(t *testing.T) {
	var cases = []struct {
		in  ProgressBar
		out string
	}{
		{ProgressBar{0.0, 10, "test.txt"}, "[  0%] test.txt"},
		{ProgressBar{0.5, 10, ""}, "[ 50%]"},
		{ProgressBar{1.0, 10, ""}, "[100%]"},
		{ProgressBar{0.26, 10, "yoruichi-rocks.img"}, "[ 26%] yoruichi-rocks.img"},
	}

	for i, tt := range cases {
		out := tt.in.Render()
		if tt.out != out {
			t.Errorf("case #%d: have %q want %q", i, out, tt.out)
		}
	}

}
func TestRenderGfx(t *testing.T) {
	var cases = []struct {
		in  ProgressBar
		out string
	}{
		{ProgressBar{0.0, 10, ""}, "[            ]"},
		{ProgressBar{0.5, 10, "FMA-OST-Brother.mp3"}, "[ .....      ] FMA-OST-Brother.mp3"},
		{ProgressBar{1.0, 10, ""}, "[ .......... ]"},
		{ProgressBar{0.26, 10, "gon.txt"}, "[ ..         ] gon.txt"},
	}

	for i, tt := range cases {
		out := tt.in.RenderGfx()
		if tt.out != out {
			t.Errorf("case #%d: have %q want %q", i, out, tt.out)
		}
	}
}
