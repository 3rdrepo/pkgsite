// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"testing"
)

func TestSeriesPathForModule(t *testing.T) {
	for _, tc := range []struct {
		modulePath, wantSeriesPath string
	}{
		{
			modulePath:     "github.com/foo",
			wantSeriesPath: "github.com/foo",
		},
		{
			modulePath:     "github.com/foo/v2",
			wantSeriesPath: "github.com/foo",
		},
		{
			modulePath:     "std",
			wantSeriesPath: "std",
		},
		{
			modulePath:     "gopkg.in/russross/blackfriday.v2",
			wantSeriesPath: "gopkg.in/russross/blackfriday",
		},
	} {
		t.Run(tc.modulePath, func(t *testing.T) {
			if got := SeriesPathForModule(tc.modulePath); got != tc.wantSeriesPath {
				t.Errorf("SeriesPathForModule(%q) = %q; want = %q", tc.modulePath, got, tc.wantSeriesPath)
			}
		})
	}
}

func TestV1Path(t *testing.T) {
	for _, test := range []struct {
		modulePath, suffix string
		want               string
	}{
		{"mod.com/foo", "bar", "mod.com/foo/bar"},
		{"mod.com/foo/v2", "bar", "mod.com/foo/bar"},
		{"std", "bar/baz", "bar/baz"},
	} {
		got := V1Path(test.modulePath, test.suffix)
		if got != test.want {
			t.Errorf("V1Path(%q, %q) = %q, want %q",
				test.modulePath, test.suffix, got, test.want)
		}
	}
}
