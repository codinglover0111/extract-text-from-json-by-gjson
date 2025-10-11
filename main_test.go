package main

import (
	"path/filepath"
	"testing"
)

func TestBuildOutputPath(t *testing.T) {
	cases := []struct {
		name      string
		inputPath string
		rex       string
		want      string
	}{
		{
			name:      "posix path",
			inputPath: "./data/sample.json",
			rex:       "payload.values",
			want:      filepath.Join(".", "sample_values.txt"),
		},
		{
			name:      "windows path",
			inputPath: `C:\\Users\\tester\\logs\\event.json`,
			rex:       "data.result",
			want:      filepath.Join(".", "event_result.txt"),
		},
		{
			name:      "quoted path with upper extension",
			inputPath: `"/tmp/Example.JSON"`,
			rex:       "root.entry",
			want:      filepath.Join(".", "Example_entry.txt"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildOutputPath(tc.inputPath, tc.rex)
			if got != tc.want {
				t.Fatalf("buildOutputPath(%q, %q) = %q, want %q", tc.inputPath, tc.rex, got, tc.want)
			}
		})
	}
}
