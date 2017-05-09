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

import "fmt"

// ProgressBar is a simple model of a progress/loading bar.
type ProgressBar struct {
	// Progress represents the current progress percentage.
	Progress float64
	// Divisions represents the granularity of the bar.
	Divisions int
	// Title is an optional title for the bar.
	Title string
}

// Render constructs a string representing a frame of the progress bar at it's
// current state.
func (p *ProgressBar) Render() string {
	out := fmt.Sprintf("[%3.f%%]", p.Progress*100)
	if p.Title != "" {
		out += " " + p.Title
	}
	return out
}

// RenderGfx is a slightly more fancy version of Render; it constructs a string
// representing the current frame of the progress bar, and attempts to emulate a
// real GUI bar's progress with text-based graphics.
func (p *ProgressBar) RenderGfx() string {
	out := "[ "
	dots := int(p.Progress * float64(p.Divisions))
	spaces := p.Divisions - dots
	for x := 0; x < dots; x++ {
		out += "."
	}
	for x := 0; x < spaces; x++ {
		out += " "
	}
	out += " ]"

	if p.Title != "" {
		out += " " + p.Title
	}

	return out
}

func (p *ProgressBar) String() string {
	return p.Render()
}
